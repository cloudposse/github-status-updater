FROM golang:1.9.3 as builder
RUN mkdir -p /go/src/github.com/cloudposse/github-status-updater
WORKDIR /go/src/github.com/cloudposse/github-status-updater
COPY . .
RUN go get && CGO_ENABLED=0 go build -v -o "./dist/bin/github-status-updater" *.go


FROM alpine:3.6
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/cloudposse/github-status-updater/dist/bin/github-status-updater /usr/bin/github-status-updater
ENV PATH $PATH:/usr/bin
ENTRYPOINT ["github-status-updater"]
