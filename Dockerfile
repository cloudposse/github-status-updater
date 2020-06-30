FROM golang:1.13.3-buster as builder
#RUN mkdir -p /go/src/github.com/cloudposse/github-status-updater
#WORKDIR /go/src/github.com/cloudposse/github-status-updater
#COPY . .
WORKDIR /usr/src
COPY . /usr/src
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -v -o "./dist/bin/github-status-updater" *.go

FROM alpine:3.12
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/src/github-status-updater /usr/bin/github-status-updater
ENV PATH $PATH:/usr/bin
ENTRYPOINT ["github-status-updater"]
