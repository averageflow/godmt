#!/bin/sh

cd ../cmd/godmt

# Linux build
GOOS=linux GOARCH=amd64 go build -o godmt-linux-x86

# macOS build
GOOS=darwin GOARCH=amd64 go build -o godmt-darwin-x86

# FreeBSD build
GOOS=freebsd GOARCH=amd64 go build -o godmt-freebsd-x86


