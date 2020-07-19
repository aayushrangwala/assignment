# Build the cipher service binary
FROM alpine:latest

# install go 1.12.7
RUN apk add --no-cache \
		ca-certificates

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

# export the golang version
ENV GOLANG_VERSION 1.12.6

# install the specific go version 1.12.6 and its dependecies
RUN apk update && apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go && \
    wget -O go.tgz https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz && \
    tar -C /usr/local -xzf go.tgz && \
    cd /usr/local/go/src/ && \
    ./make.bash && \
    export PATH="/usr/local/go/bin:$PATH" && \
    export GOPATH=/go && \
    export PATH=$PATH:$GOPATH/bin && \
    apk del .build-deps && \
    go version

# Export env var for GOPATH and set it under the PATH
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# make the GOPATH and bin directory
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# add the working directory for the project
WORKDIR "$GOPATH/src/discovergy"

# just an indication that this port will be exposed by this container
EXPOSE 3333

# Copy the service code
COPY internal internal
COPY pkg pkg
COPY vendor vendor
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o www

# command to run at the immediate start of the container
ENTRYPOINT ["./www"]
