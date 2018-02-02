# github-commit-status

Command line utility for setting GitHub commit status.
Useful for CI environments like Travis, Circle or CodeFresh to set more specific commit statuses on pull requests.
It accepts parameters as command-line arguments or as ENV variables.

__NOTE__: Create a [GitHub token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) with `repo:status` scope


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
export GITHUB_COMMIT_TARGET_URL=https://ci.example.com/build/1

github-commit-status
```



### Run locally with command-line parameters

```sh
go get

CGO_ENABLED=0 go build -v -o "./github-commit-status" *.go

github-commit-status -token XXXXXXXXXXXXXXXX \
                     -owner cloudposse \
                     -repo github-commit-status \
                     -sha XXXXXXXXXXXXXXX \
                     -state success \
                     -context CI \
                     -description "Commit status with target URL" \
                     -url https://ci.example.com/build/1
```



### Run in a Docker container with ENV vars
__NOTE__: it will download all `Go` dependencies and build the program inside the container


```sh
docker run --rm github-commit-status -e "GITHUB_TOKEN=XXXXXXXXXXXXXXXX" \
                                     -e "GITHUB_OWNER=cloudposse" \
                                     -e "GITHUB_REPO=github-commit-status" \
                                     -e "GITHUB_COMMIT_SHA=XXXXXXXXXXXXXXXX" \
                                     -e "GITHUB_COMMIT_STATE=success" \
                                     -e "GITHUB_COMMIT_CONTEXT=CI" \
                                     -e "GITHUB_COMMIT_DESCRIPTION=Commit status with target URL" \
                                     -e "GITHUB_COMMIT_TARGET_URL=https://ci.example.com/build/1"
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
export GITHUB_COMMIT_TARGET_URL=https://ci.example.com/build/1

docker run --rm github-commit-status -e GITHUB_TOKEN \
                                     -e GITHUB_OWNER \
                                     -e GITHUB_REPO \
                                     -e GITHUB_COMMIT_SHA \
                                     -e GITHUB_COMMIT_STATE \
                                     -e GITHUB_COMMIT_CONTEXT \
                                     -e GITHUB_COMMIT_DESCRIPTION \
                                     -e GITHUB_COMMIT_TARGET_URL
```



## References
* https://github.com/google/go-github/
* https://docs.docker.com/engine/reference/commandline/run/
* https://docs.docker.com/develop/develop-images/dockerfile_best-practices/



## LICENSE
See [LICENSE](LICENSE)
