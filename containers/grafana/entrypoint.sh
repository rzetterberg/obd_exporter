# Entrypoint that allows us to automate the creation of the Prometheus Data
# Source. This is the "best" way currently, a issue is open on GitHub to
# add a feature that allows users to do this more elegantly.

ARGS="$@"

/run.sh "$ARGS" &
PID=$!

AddDataSource() {
    curl 'http://admin:dev@localhost:3000/api/datasources' \
         -X POST \
         -H 'Content-Type: application/json;charset=UTF-8' \
         --data-binary \
         '{"name":"Default","type":"prometheus","url":"http://prometheus:9090","access":"proxy","isDefault":true}'
}

until AddDataSource; do
    sleep 1
done

kill $PID

exec /bin/sh /run.sh "$ARGS"
