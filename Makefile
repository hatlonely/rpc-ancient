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
	@echo "install golang dependency"
	go mod tidy
	go mod vendor

codegen: api/ancient.proto
	mkdir -p api/gen/go && mkdir -p api/gen/swagger
	protoc -I.. -I. --gofast_out=plugins=grpc,paths=source_relative:api/gen/go/ $<
	protoc -I.. -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:api/gen/go $<
	protoc -I.. -I. --swagger_out=logtostderr=true:api/gen/swagger $<

.PHONY: image
image:
	docker build --tag=hatlonely/${repository}:${version} .
