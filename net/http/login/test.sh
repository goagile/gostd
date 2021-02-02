#!/bin/bash

echo "Test login form"

go run login.go &
google-chrome-stable http://127.0.0.1:8080
