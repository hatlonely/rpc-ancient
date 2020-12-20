binary=ancient
dockeruser=hatlonely
gituser=hatlonely
repository=go-rpc-ancient
version=1.0.0
export GOPROXY=https://goproxy.cn

.PHONY: build
build: cmd/main.go Makefile vendor
	mkdir -p build/bin
	go build -ldflags "-X 'main.Version=`sh scripts/version.sh`'" cmd/main.go && mv main build/bin/${binary} && cp -r config build/

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
