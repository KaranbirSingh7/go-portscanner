.PHONY: help

help:
	./portscan -h

run: 
	echo "Running main.go"
	go run main.go

build:
	echo "Building go binary"
	go build -o portscan main.go