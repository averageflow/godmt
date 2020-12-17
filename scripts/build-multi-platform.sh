#!/bin/sh

cd ../cmd/goschemaconverter

# Linux build
GOOS=linux GOARCH=amd64 go build -o goschemaconverter-linux-x86

# macOS build
GOOS=darwin GOARCH=amd64 go build -o goschemaconverter-darwin-x86

# FreeBSD build
GOOS=freebsd GOARCH=amd64 go build -o goschemaconverter-freebsd-x86


