#!/bin/bash

echo "Test Upload Raw JSON"

curl -X POST \
    -H "Content-Type:application/json" \
    -d '{"name":"Go"}' \
    127.0.0.1:8080/raw
