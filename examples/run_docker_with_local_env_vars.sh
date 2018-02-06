#!/bin/bash

export GITHUB_ACTION=update_state
export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
export GITHUB_OWNER=cloudposse
export GITHUB_REPO=github-commit-status
export GITHUB_REF=XXXXXXXXXXXXXXXX
export GITHUB_STATE=success
export GITHUB_CONTEXT="my-ci"
export GITHUB_DESCRIPTION="Commit status with target URL"
export GITHUB_TARGET_URL="https://my-ci.com/build/1"

docker run -i --rm \
        -e GITHUB_ACTION \
        -e GITHUB_TOKEN \
        -e GITHUB_OWNER \
        -e GITHUB_REPO \
        -e GITHUB_REF \
        -e GITHUB_STATE \
        -e GITHUB_CONTEXT \
        -e GITHUB_DESCRIPTION \
        -e GITHUB_TARGET_URL \
        github-commit-status
