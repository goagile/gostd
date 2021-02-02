#!/bin/bash

echo "Test Query param and Form Key"

curl "127.0.0.1:8080?param=111"

curl "127.0.0.1:8080?key=222"
