#!/usr/bin/env bash

CGO_ENABLED=0  GOOS=darwin GOARCH=amd64  go build -o build/darwin-amd64/xyz ./cmd
CGO_ENABLED=0  GOOS=darwin GOARCH=arm64  go build -o build/darwin-arm64/xyz ./cmd
