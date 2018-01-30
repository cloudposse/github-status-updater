#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

GITHUB_USER=thbishop
PROJECT_NAME=github-commit-status

create_release() {
    echo ""
    echo "Creating release '$version'."

    env GITHUB_API=https://api.github.com \
        GITHUB_REPO=$PROJECT_NAME \
        github-release release --tag $version --name $version --user $GITHUB_USER

    echo ""
    echo "Release '$version' created."
}

create_tag() {
    git tag -m "$version" $version
    echo "Created tag '$version'"
}

ensure_clean_and_committed() {
    set +e
    if ! git diff --exit-code --quiet || ! git diff --cached --exit-code --quiet; then
        echo "There are files that need to be committed or removed first."
        echo ""
        echo "$(git status)"
        exit 1
    fi
    set -e
}

ensure_on_master() {
    branch=$(git symbolic-ref --short -q HEAD)
    if [[ "$branch" != "master" ]]; then
        echo "Must be on master to create a release."
        echo "Currently on '$branch'."
        exit 1
    fi
}

ensure_release_cli_exists() {
    if [ ! $(which github-release) ]; then
        echo "Unable to find 'github-release' in your \$PATH."
        echo "If you're OSX you can 'brew install github-release'."
        echo "Otherwise, see https://github.com/aktau/github-release"
        exit 1
    fi
}

ensure_token() {
    set +u
    if [ -z "$GITHUB_TOKEN" ]; then
        echo "Env var 'GITHUB_TOKEN' is not set"
        exit 1
    fi
    set -u
}

push_to_origin() {
    echo "Pushing master and tags to origin"
    git push
    git push --tags
}

upload_artifacts() {
    echo ""
    upload_url="$(curl -s -H 'Accept: application/json' \
                 https://api.github.com/repos/$GITHUB_USER/$PROJECT_NAME/releases | \
                 jq  -M .[].upload_url | sed -E 's/{.+$//g' | sed -E 's/"//g')"

    for f in $(ls pkg/*{.rpm,.tar.gz,sums}); do
        echo "Uploading '$(basename $f)' for release '$version'"
        curl -s \
             -H "Authorization: token $GITHUB_TOKEN" \
             -H "Accept: application/vnd.github.v3+json" \
             -H "Content-Type: application/octet-stream" \
             --data-binary @$f \
             "$upload_url?name=$(basename $f)" > /dev/null
    done
}

ensure_token
ensure_release_cli_exists
ensure_on_master
ensure_clean_and_committed

version=$(grep version $PROJECT_NAME/version.go |
          awk '{print $4}' |
          sed 's/"//g')

create_tag
push_to_origin
create_release
upload_artifacts

echo "Release complete!"
