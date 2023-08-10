###################################################
FROM ubuntu:latest AS certs
RUN apt update
RUN apt install -y ca-certificates

##################################################
FROM golang:1.19-alpine as builder

ARG BUILD_NUM=1
ARG BUILD_USER=jenkins
ARG GOPROXY=

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git make musl-dev curl busybox-static jq bash tree wget

WORKDIR /go/src/github.com/af-go/basic-app
COPY go.mod .
COPY go.sum .

COPY Makefile .
RUN make go-download && make dep
COPY . .

# Run all test and build steps one by one
# RUN make check-fmt 
# RUN make lint 
RUN make build

RUN wget https://rolesanywhere.amazonaws.com/releases/1.0.5/X86_64/Linux/aws_signing_helper

RUN mv aws_signing_helper /go/src/github.com/af-go/basic-app/dist/


############################
#  build image
############################
FROM ubuntu:latest

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder  /go/src/github.com/af-go/basic-app/dist/basic-app /usr/local/bin/basic-app

COPY --from=builder /go/src/github.com/af-go/basic-app/dist/aws_signing_helper /usr/local/bin/aws_signing_helper

RUN chmod 0755 /usr/local/bin/aws_signing_helper



# Run the binary.
ENTRYPOINT ["/usr/local/bin/basic-app"]

EXPOSE 8080

CMD ["agent"]
