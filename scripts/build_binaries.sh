#!/bin/bash

set -eo pipefail

OS_TYPES=(darwin linux)
PROJECT_ROOT="$( cd -P "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
PKG_ROOT="$PROJECT_ROOT/pkg"
PROJECT_NAME="github-commit-status"
BIN_NAME="github-commit-status"

echo -n "Removing existing binaries..."
rm -fr $PKG_ROOT
echo "Done"

mkdir -p $PKG_ROOT

version=$(grep version $PROJECT_NAME/version.go  | awk '{print $4}' | sed 's/"//g')

echo "Building packages for distribution for version $version."

for os in ${OS_TYPES[*]}; do
    pkg_name="$BIN_NAME-$version-$os-amd64"
    pkg_path="$PKG_ROOT/$pkg_name"
    echo -n "Building binary for $os/amd64..."
    env GOOS=$os GOARCH=amd64 go build -o $pkg_path/$BIN_NAME ./$PROJECT_NAME
    cd $PKG_ROOT
    tar zcf $pkg_name.tar.gz -C $pkg_path $BIN_NAME
    shasum --algorithm 256 --binary $pkg_name.tar.gz >> $version-sha256-sums
    cd $PROJECT_ROOT
    echo "Done"
done

echo ""
echo "Done"
