FROM golang:1.8.3

WORKDIR /go/src/odb_exporter
COPY src/github.com/rzetterberg/obd_exporter .

RUN go-wrapper download && \
    go-wrapper install

ENTRYPOINT ["go-wrapper", "run"]
