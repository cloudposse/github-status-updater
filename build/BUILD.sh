#!/bin/sh

echo "Building github-commit-status"
CGO_ENABLED=0 go build -v -o "./dist/bin/github-commit-status" *.go
