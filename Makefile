SHELL = /bin/bash

PATH:=$(PATH):$(GOPATH)/bin

include $(shell curl --silent -o .build-harness "https://raw.githubusercontent.com/cloudposse/build-harness/master/templates/Makefile.build-harness"; echo .build-harness)


.PHONY : go-get
go-get:
	go get


.PHONY : go-build
go-build: go-get
	CGO-ENABLED=0 go build -v -o "./dist/bin/github-commit-status" *.go


.PHONY : export-env
export-env:
	export GITHUB-TOKEN=XXXXXXXXXXXXXXXX
	export GITHUB-OWNER=cloudposse
	export GITHUB-REPO=github-commit-status
	export GITHUB-COMMIT-SHA=XXXXXXXXXXXXXXXX
	export GITHUB-COMMIT-STATE=success
	export GITHUB-COMMIT-CONTEXT=CI
	export GITHUB-COMMIT-DESCRIPTION="Commit status with target URL"
	export GITHUB-COMMIT-TARGET-URL=https://my.buildstatus.com/build/3


.PHONY : run-locally-with-env-vars
run-locally-with-env-vars: go-build export-env
	./dist/bin/github-commit-status


.PHONY : run-locally-with-command-line-args
run-locally-with-command-line-args: go-build
	./dist/bin/github-commit-status \
            -token XXXXXXXXXXXXXXXX \
            -owner cloudposse \
            -repo github-commit-status \
            -sha XXXXXXXXXXXXXXX \
            -state success \
            -context CI \
            -description "Commit status with target URL" \
            -url https://my.buildstatus.com/build/3


.PHONY : docker-build
docker-build:
	docker build --tag github-commit-status  --no-cache=true .


.PHONY : run-docker-with-env-vars
run-docker-with-env-vars: docker-build
	docker run -it --rm \
            -e GITHUB-TOKEN=XXXXXXXXXXXXXXXX \
            -e GITHUB-OWNER=cloudposse \
            -e GITHUB-REPO=github-commit-status \
            -e GITHUB-COMMIT-SHA=XXXXXXXXXXXXXXXX \
            -e GITHUB-COMMIT-STATE=success \
            -e GITHUB-COMMIT-CONTEXT=CI \
            -e GITHUB-COMMIT-DESCRIPTION="Commit status with target URL" \
            -e GITHUB-COMMIT-TARGET-URL=https://my.buildstatus.com/build/3 \
            github-commit-status


.PHONY : run-docker-with-local-env-vars
run-docker-with-local-env-vars: docker-build export-env
	docker run -it --rm \
            -e GITHUB-TOKEN \
            -e GITHUB-OWNER \
            -e GITHUB-REPO \
            -e GITHUB-COMMIT-SHA \
            -e GITHUB-COMMIT-STATE \
            -e GITHUB-COMMIT-CONTEXT \
            -e GITHUB-COMMIT-DESCRIPTION \
            -e GITHUB-COMMIT-TARGET-URL \
            github-commit-status


.PHONY : run-docker-with-env-vars-file
run-docker-with-env-vars-file: docker-build
	docker run -it --rm --env-file ./env.list github-commit-status
