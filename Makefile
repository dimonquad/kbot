APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=dimonquad
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
GOARCH=arm64

format:
	gofmt -s -w ./

lint:
	golint

get:
    go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${GOARCH} go build -v -o kbot -ldflags "-X="github.com/dimonquad/kbot/cmd.appVersion=${VERSION}

image:
	docker build -t ${REGISTRY}/${APP}:${VERSION}-${GOARCH}.

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${GOARCH}

clean:
	rm -rf kbot 