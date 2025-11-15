#APP=$("kbot" $(shell git remote get-url origin))
APP=kbot
REGISTRY=dimonquad
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
GOARCH=arm64

format:
	gofmt -s -w ./

get:
	go get

#ifeq ($(OS),Windows_NT)
build: format get 
	powershell -NoProfile -Command "$$env:CGO_ENABLED='0'; $$env:GOOS='$(TARGETOS)'; $$env:GOARCH='$(GOARCH)'; go build -v -o kbot -ldflags \"-X=github.com/dimonquad/kbot/cmd.appVersion=$(VERSION)\""
#else
#build: format get
    #CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${GOARCH} go build -v -o kbot -ldflags "-X=github.com/dimonquad/kbot/cmd.appVersion=${VERSION}"
#endif

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${GOARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${GOARCH}

clean:
	rm -rf kbot 