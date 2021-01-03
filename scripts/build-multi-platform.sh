#!/bin/sh

cd ../cmd/godmt

# Linux build
GOOS=linux GOARCH=amd64 go build -o godmt-linux-amd64

# macOS build
GOOS=darwin GOARCH=amd64 go build -o godmt-darwin-amd64

# FreeBSD build
GOOS=freebsd GOARCH=amd64 go build -o godmt-freebsd-amd64

# Windows build
GOOS=windows GOARCH=amd64 go build -o godmt-windows-amd64


