#!/bin/bash

export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
export GITHUB_OWNER=cloudposse
export GITHUB_REPO=github-commit-status
export GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX
export GITHUB_COMMIT_STATE=success
export GITHUB_COMMIT_CONTEXT=CI
export GITHUB_COMMIT_DESCRIPTION="Commit status with target URL"
export GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3

./dist/bin/github-commit-status
