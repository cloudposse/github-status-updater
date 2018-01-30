#!/usr/bin/env bash
# chmod 777 build_linux_docker.sh

_tag=$1

if [ -z "${_tag}" ]; then
    source ./_VERSION
    _tag=${_VERSION}
fi

echo 'Building Go application'
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GOROOT=/usr/local/go go build -a -tags netgo -ldflags '-w'

echo 'Building Docker Image'
docker build --tag "cloudposse/github-commit-status:${_tag}"  --no-cache=true .
