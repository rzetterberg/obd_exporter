package main

import (
	"github.com/rzetterberg/elmobd"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Output map[string]string

// Friendly reminder for anyone reading this code, be careful when using
// this value. Storing a nil value or a value of another type than the
// first value will result in a panic.
var GlobalOutput atomic.Value

func newOBDSession(serialDevice string) error {
	dev, err := elmobd.NewDevice(serialDevice, false)

	if err != nil {
		return err
	}

	supported, err := dev.CheckSupportedCommands()

	if err != nil {
		return err
	}

	supportedCmds := supported.FilterSupported(
		elmobd.GetSensorCommands(),
	)

	for {
		results, err := dev.RunManyOBDCommands(supportedCmds)

		if err != nil {
			return err
		}

		output := make(Output)

		for _, res := range results {
			output[res.Key()] = res.ValueAsLit()
		}

		GlobalOutput.Store(output)

		time.Sleep(time.Millisecond * 10)
	}
}

func startOBDReading(serialDevice string) {
	sleepAmount := 1

	for {
		log.Println("Starting new OBD session")
		err := newOBDSession(serialDevice)

		log.Println("OBD session closed by error: ", err)

		GlobalOutput.Store(make(Output))

		sleepAmount *= 2

		log.Printf("Waiting %d seconds before new OBD session\n", sleepAmount)
		time.Sleep(time.Second * time.Duration(sleepAmount))
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OBD exporter!\n")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	output := GlobalOutput.Load().(Output)

	for k, v := range output {
		name := "obd_" + k

		fmt.Fprintf(w, "# HELP %s No description\n", name)
		fmt.Fprintf(w, "# TYPE %s gauge\n", name)
		fmt.Fprintf(w, "%s %s\n\n", name, v)
	}
}

func startWebServing() {
	log.Println("Starting web serving")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/metrics", metricsHandler)

	http.ListenAndServe(":8080", nil)
}

func main() {
	var serialDevice = flag.String("serial-device", "/dev/ttyUSB0", "Serial device to use")

	flag.Parse()

	GlobalOutput.Store(
		make(Output),
	)

	go startOBDReading(*serialDevice)

	startWebServing()
}
