SHELL = /bin/bash

PATH:=$(PATH):$(GOPATH)/bin

include $(shell curl --silent -o .build-harness "https://raw.githubusercontent.com/cloudposse/build-harness/master/templates/Makefile.build-harness"; echo .build-harness)


.PHONY : go-get
go-get:
	go get


.PHONY : go-build
go-build: go-get
	CGO_ENABLED=0 go build -v -o "./dist/bin/github-commit-status" *.go


.PHONY : run-locally-with-env-vars
run-locally-with-env-vars: go-build
	./run_locally_with_env_vars.sh


.PHONY : run-locally-with-command-line-args
run-locally-with-command-line-args: go-build
	./run_locally_with_command_line_args.sh


.PHONY : docker-build
docker-build:
	docker build --tag github-commit-status  --no-cache=true .


.PHONY : run-docker-with-env-vars
run-docker-with-env-vars: docker-build
	./run_docker_with_env_vars.sh


.PHONY : run-docker-with-local-env-vars
run-docker-with-local-env-vars: docker-build
	./run_docker_with_local_env_vars.sh


.PHONY : run-docker-with-env-vars-file
run-docker-with-env-vars-file: docker-build
	./run_docker_with_env_vars_file.sh
