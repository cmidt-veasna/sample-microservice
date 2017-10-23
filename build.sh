#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -v -o service main.go

docker build -t sample2:v1.0.0 .

# kubectl set image deployment/sample1 sample1=sample1:v1.0.0