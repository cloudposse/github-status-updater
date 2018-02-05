#!/bin/bash

docker run -i --rm \
        -e GITHUB_TOKEN=XXXXXXXXXXXXXXXX \
        -e GITHUB_OWNER=cloudposse \
        -e GITHUB_REPO=github-commit-status \
        -e GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX \
        -e GITHUB_COMMIT_STATE=success \
        -e GITHUB_COMMIT_CONTEXT=CI \
        -e GITHUB_COMMIT_DESCRIPTION="Commit status with target URL" \
        -e GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3 \
        github-commit-status
