binary=ancient
repository=rpc-ancient
version=$(shell git describe --tags | awk '{print(substr($$0,2,length($$0)))}')
export GOPROXY=https://goproxy.cn

define BUILD_VERSION
  version: $(shell git describe --tags)
gitremote: $(shell git remote -v | grep fetch | awk '{print $$2}')
   commit: $(shell git rev-parse HEAD)
 datetime: $(shell date '+%Y-%m-%d %H:%M:%S')
 hostname: $(shell hostname):$(shell pwd)
goversion: $(shell go version)
endef
export BUILD_VERSION

.PHONY: build
build: cmd/main.go Makefile vendor
	mkdir -p build/bin
	go build -ldflags "-X 'main.Version=$$BUILD_VERSION'" cmd/main.go && mv main build/bin/${binary} && cp -r config build/

vendor: go.mod go.sum
	go mod tidy
	go mod vendor

.PHONY: codegen
codegen: api/ancient.proto submodule
	mkdir -p api/gen/go && mkdir -p api/gen/swagger
	protoc -Irpc-api -I. --gofast_out=plugins=grpc,paths=source_relative:api/gen/go/ $<
	protoc -Irpc-api -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:api/gen/go $<
	protoc -Irpc-api -I. --swagger_out=logtostderr=true:api/gen/swagger $<

.PHONY: submodule
submodule:
	 git submodule update

.PHONY: image
image:
	docker build --tag=hatlonely/${repository}:${version} .
