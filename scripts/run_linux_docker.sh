#!/usr/bin/env bash
# chmod 777 run_linux_docker.sh

_tag=$1

if [ -z "${_tag}" ]; then
    source _VERSION
    _tag=${_VERSION}
fi

echo 'Running Docker Image'
docker run --env-file ./env.list -p 5000:5000 -d "cloudposse/github-commit-status:${_tag}"
