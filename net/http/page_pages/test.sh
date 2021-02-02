#!/bin/bash

echo "Test HTTP server"
echo "Test Home request"
curl 127.0.0.1:8080

echo "Test Page request"
curl 127.0.0.1:8080/page

echo "Test Fail -> Page as dir"
curl 127.0.0.1:8080/page/xyz

echo "Test Pages/XYZ ..."
curl 127.0.0.1:8080/pages/xyz
curl 127.0.0.1:8080/pages/abc
