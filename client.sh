#!/bin/bash

while true; do
    payload='{"name":"testClient","email":"testing@middleware.si"}'
    echo "PUT: $payload"
    echo "response: "
    curl -X PUT -H "Content-Type: application/json" -d $payload localhost:8080/hello
    echo
    echo
    sleep 2
done
