#!/usr/bin/env bash

CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o build/linux-amd64/xyz ./cmd
CGO_ENABLED=0  GOOS=linux GOARCH=arm64 go build -o build/linux-arm64/xyz ./cmd
