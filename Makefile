SHELL = /bin/bash

.PHONY : go_get
go_get:
	go get


.PHONY : go_build
go_build: go_get
	CGO_ENABLED=0 go build -v -o "./dist/bin/github-commit-status" *.go


.PHONY : export_env
export_env:
	export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
	export GITHUB_OWNER=cloudposse
	export GITHUB_REPO=github-commit-status
	export GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX
	export GITHUB_COMMIT_STATE=success
	export GITHUB_COMMIT_CONTEXT=CI
	export GITHUB_COMMIT_DESCRIPTION="Commit status with target URL"
	export GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3


.PHONY : run_locally_with_env_vars
run_locally_with_env_vars: go_build export_env
	./dist/bin/github-commit-status


.PHONY : run_locally_with_command_line_args
run_locally_with_command_line_args: go_build
	./dist/bin/github-commit-status \
            -token XXXXXXXXXXXXXXXX \
            -owner cloudposse \
            -repo github-commit-status \
            -sha XXXXXXXXXXXXXXX \
            -state success \
            -context CI \
            -description "Commit status with target URL" \
            -url https://my.buildstatus.com/build/3


.PHONY : docker_build
docker_build:
	docker build --tag github-commit-status  --no-cache=true .


.PHONY : run_docker_with_env_vars
run_docker_with_env_vars: docker_build
	docker run -it --rm \
            -e GITHUB_TOKEN=XXXXXXXXXXXXXXXX \
            -e GITHUB_OWNER=cloudposse \
            -e GITHUB_REPO=github-commit-status \
            -e GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX \
            -e GITHUB_COMMIT_STATE=success \
            -e GITHUB_COMMIT_CONTEXT=CI \
            -e GITHUB_COMMIT_DESCRIPTION="Commit status with target URL" \
            -e GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3 \
            github-commit-status


.PHONY : run_docker_with_local_env_vars
run_docker_with_local_env_vars: docker_build export_env
	docker run -it --rm \
            -e GITHUB_TOKEN \
            -e GITHUB_OWNER \
            -e GITHUB_REPO \
            -e GITHUB_COMMIT_SHA \
            -e GITHUB_COMMIT_STATE \
            -e GITHUB_COMMIT_CONTEXT \
            -e GITHUB_COMMIT_DESCRIPTION \
            -e GITHUB_COMMIT_TARGET_URL \
            github-commit-status


.PHONY : run_docker_with_env_vars_file
run_docker_with_env_vars_file: docker_build
	docker run -it --rm --env-file ./env.list github-commit-status
