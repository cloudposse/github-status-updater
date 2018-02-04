FROM golang:1.9.3 as builder
RUN mkdir -p /go/src/github.com/cloudposse/github-commit-status
WORKDIR /go/src/github.com/cloudposse/github-commit-status
COPY . .
RUN go get && CGO_ENABLED=0 go build -v -o "./dist/bin/github-commit-status" *.go


FROM alpine:3.6
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/cloudposse/github-commit-status/dist/bin/github-commit-status /usr/bin/github-commit-status
ENV PATH $PATH:/usr/bin
ENTRYPOINT ["github-commit-status"]
