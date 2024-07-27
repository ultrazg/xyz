#!/usr/bin/env bash

CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/xyz-linux-amd64 .
CGO_ENABLED=0  GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/xyz-inux-arm64 .
CGO_ENABLED=0  GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o build/xyz-darwin-amd64 .
CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o build/xyz-darwin-arm64 .
CGO_ENABLED=0  GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o build/xyz-windows-amd64.exe .
