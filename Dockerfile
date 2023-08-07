############################
# STEP 1 build executable binary
############################
FROM golang:1.19-alpine as builder

ARG BUILD_NUM=1
ARG BUILD_USER=jenkins
ARG GOPROXY=

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git make musl-dev curl busybox-static jq bash tree

WORKDIR /go/src/github.com/af-go/basic-app
COPY go.mod .
COPY go.sum .

COPY Makefile .
RUN make go-download && make dep
COPY . .

# Run all test and build steps one by one
RUN make check-fmt 
RUN make lint 
RUN make build

COPY thirdparty/aws_signing_helper /go/src/github.com/af-go/basic-app/dist/aws_signing_helper
############################
# STEP 2 build a small image
############################
FROM alpine:3.12

RUN apk update && apk add curl

COPY --from=builder  /go/src/github.com/af-go/basic-app/dist/basic-app /basic-app

COPY --from=builder /go/src/github.com/af-go/basic-app/dist/aws_signing_helper /aws_signing_helper

# Run the binary.
ENTRYPOINT ["/basic-app"]

EXPOSE 8080

CMD ["agent"]
