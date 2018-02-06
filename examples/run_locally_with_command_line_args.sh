#!/bin/bash

../dist/bin/github-commit-status \
        -action update_state \
        -token XXXXXXXXXXXXXXXX \
        -owner cloudposse \
        -repo github-commit-status \
        -ref XXXXXXXXXXXXXXX \
        -state success \
        -context "my-ci" \
        -description "Commit status with target URL" \
        -url "https://my-ci.com/build/1"
