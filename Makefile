all: clean test build

binaries: clean fmt test
	@script/build_binaries.sh

build:
	@echo "==> Compiling source code."
	@godep go build -v -o ./bin/github-commit-status ./github-commit-status

clean:
	@echo "==> Cleaning up previous builds."
	@rm -rf bin/github-commit-status

deps:
	@echo "==> Downloading dependencies."
	@godep save -r ./github-commit-status/...

fmt:
	@echo "==> Formatting source code."
	@goimports -w ./github-commit-status

release:
	@echo "==> Releasing"
	@script/release

test: fmt vet
	@echo "==> Running tests."
	@godep go test -cover ./github-commit-status/...

vet:
	@godep go vet ./github-commit-status/...

help:
	@echo "binaries\tcreate binaries"
	@echo "build\t\tbuild the code"
	@echo "clean\t\tremove previous builds"
	@echo "deps\t\tdownload dependencies"
	@echo "fmt\t\tformat the code"
	@echo "release\t\tcreate a release"
	@echo "test\t\ttest the code"
	@echo "vet\t\tvet the code"
	@echo ""
	@echo "default will test, format, and build the code"

.PNONY: all clean deps fmt help test
