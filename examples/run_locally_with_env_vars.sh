#!/bin/bash

export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
export GITHUB_OWNER=cloudposse
export GITHUB_REPO=github-commit-status
export GITHUB_REF=XXXXXXXXXXXXXXXX
export GITHUB_STATE=success
export GITHUB_CONTEXT=CI
export GITHUB_DESCRIPTION="Commit status with target URL"
export GITHUB_TARGET_URL=https://my.buildstatus.com/build/3

../dist/bin/github-commit-status
