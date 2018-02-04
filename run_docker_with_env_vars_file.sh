#!/bin/bash

docker run -i --rm --env-file ./env.list github-commit-status
