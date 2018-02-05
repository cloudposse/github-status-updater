#!/bin/bash

../dist/bin/github-commit-status \
        -token XXXXXXXXXXXXXXXX \
        -owner cloudposse \
        -repo github-commit-status \
        -sha XXXXXXXXXXXXXXX \
        -state success \
        -context CI \
        -description "Commit status with target URL" \
        -url https://my.buildstatus.com/build/3
