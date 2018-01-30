# github-commit-status

CLI to update the status of a GitHub commit


This is a simple utility to update the status of a commit on github. The
primary use case is to update the status of a commit in a build environment.


## Install

Download the latest binary or
`brew tap thbishop/github-commit-status && brew install github-commit-status`
if you're on OSX.


## Usage

Create a [github token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/)
with `repo:status` scope and export it as an env var:
```sh
export GITHUB_TOKEN=1234
```

Update the status with:
```sh
github-commit-status --user foo --repo bar --commit $SHA --state success
```

You can also optionally include a target url, description, or context to be
included in the status update:
```sh
github-commit-status --user foo \
                     --repo bar \
                     --commit $SHA \
                     --state success --target-url https://ci.example.com/build/1 \
                     --description "It failed because it is busted" \
                     --context ci
```

If you're using github enterprise, you can set the API endpoint with an env
var like so:
```sh
export GITHUB_API=https://github.example.com/api/v3
```

If needed, a proxy can be configured using environment variables:
* `http_proxy`
* `HTTP_PROXY`


## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Make sure tests pass
* Send a pull request and I'll get it integrated


## LICENSE
See [LICENSE](LICENSE)
