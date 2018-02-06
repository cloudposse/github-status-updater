#!/bin/bash

docker run -i --rm \
        -e GITHUB_ACTION=update_state \
        -e GITHUB_TOKEN=XXXXXXXXXXXXXXXX \
        -e GITHUB_OWNER=cloudposse \
        -e GITHUB_REPO=github-commit-status \
        -e GITHUB_REF=XXXXXXXXXXXXXXXX \
        -e GITHUB_STATE=success \
        -e GITHUB_CONTEXT="my-ci" \
        -e GITHUB_DESCRIPTION="Commit status with target URL" \
        -e GITHUB_TARGET_URL="https://my-ci.com/build/1" \
        github-commit-status
