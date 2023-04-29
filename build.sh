#!/bin/bash

# set the output directory
OUT_DIR=./bin

# set the name of the program
PROGRAM_NAME=empedoc

# build for Linux
GOOS=linux GOARCH=amd64 go build -o $OUT_DIR/$PROGRAM_NAME-linux-amd64

# build for Windows
GOOS=windows GOARCH=amd64 go build -o $OUT_DIR/$PROGRAM_NAME-windows-amd64.exe
