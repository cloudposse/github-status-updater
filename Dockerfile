FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/codefresh-io/github-commit-status
WORKDIR /go/src/github.com/codefresh-io/github-commit-status
COPY . .
RUN "./build/BUILD.sh"


FROM alpine:3.6
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/codefresh-io/github-commit-status/dist/bin/github-commit-status /usr/bin/github-commit-status
ENV PATH $PATH:/usr/bin/github-commit-status
ENTRYPOINT ["github-commit-status"]
CMD ["--help"]
