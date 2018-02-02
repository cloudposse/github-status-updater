# github-commit-status

Command line utility for setting or updating GitHub commit status.

Useful for CI environments like Travis, Circle or CodeFresh to set more specific commit statuses, including setting the target URL (the URL of the page representing the status).

It accepts parameters as command-line arguments or as ENV variables.



__NOTE__: Create a [GitHub token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) with `repo:status` scope


__NOTE__: `-state` or `GITHUB_COMMIT_STATE` must be one of `error`, `failure`, `pending`, `success`



## Usage


### Run locally with ENV vars


```sh
go get

CGO_ENABLED=0 go build -v -o "./github-commit-status" *.go

export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
export GITHUB_OWNER=cloudposse
export GITHUB_REPO=github-commit-status
export GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX
export GITHUB_COMMIT_STATE=success
export GITHUB_COMMIT_CONTEXT=CI
export GITHUB_COMMIT_DESCRIPTION="Commit status with target URL"
export GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3

./github-commit-status
```



### Run locally with command-line arguments


```sh
go get

CGO_ENABLED=0 go build -v -o "./github-commit-status" *.go

./github-commit-status \
            -token XXXXXXXXXXXXXXXX \
            -owner cloudposse \
            -repo github-commit-status \
            -sha XXXXXXXXXXXXXXX \
            -state success \
            -context CI \
            -description "Commit status with target URL" \
            -url https://my.buildstatus.com/build/3
```



### Run in a Docker container with ENV vars
__NOTE__: it will download all `Go` dependencies and then build and run the program inside the container (see [`Dockerfile`](Dockerfile))


```sh
docker build --tag github-commit-status  --no-cache=true .

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
```



### Run in a Docker container with local ENV vars propagated into the container's environment


```sh
export GITHUB_TOKEN=XXXXXXXXXXXXXXXX
export GITHUB_OWNER=cloudposse
export GITHUB_REPO=github-commit-status
export GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX
export GITHUB_COMMIT_STATE=success
export GITHUB_COMMIT_CONTEXT=CI
export GITHUB_COMMIT_DESCRIPTION="Commit status with target URL"
export GITHUB_COMMIT_TARGET_URL=https://my.buildstatus.com/build/3

docker build --tag github-commit-status  --no-cache=true .

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
```



### Run in a Docker container with ENV vars declared in a file


```sh
docker build --tag github-commit-status  --no-cache=true .

docker run -it --rm --env-file ./env.list github-commit-status
```




## References
* https://github.com/google/go-github
* https://docs.docker.com/develop/develop-images/dockerfile_best-practices
* https://docs.docker.com/engine/reference/commandline/build
* https://docs.docker.com/engine/reference/commandline/run/



## LICENSE
See [LICENSE](LICENSE)
