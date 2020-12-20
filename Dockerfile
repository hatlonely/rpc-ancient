FROM golang:1.14 AS build

COPY . /go/src/
WORKDIR /go/src/
RUN make build

FROM centos:centos7
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY --from=build /go/src/build /work/go-rpc-ancient
WORKDIR /work/go-rpc-ancient
CMD [ "bin/ancient", "-c", "config/go-rpc-ancient.json" ]
