#!/bin/bash

echo "Test HTTP server"

curl 127.0.0.1:8080/
curl 127.0.0.1:8080/page
curl 127.0.0.1:8080/pages/xyz
curl 127.0.0.1:8080/pages/abc
