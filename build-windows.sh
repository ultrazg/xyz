#!/usr/bin/env bash

CGO_ENABLED=0  GOOS=windows GOARCH=amd64  go build -o build/windows-amd64/xyz.exe ./cmd
