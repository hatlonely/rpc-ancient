#!/usr/bin/env bash

export NAMESPACE="prod"
export NAME="rpc-ancient"
export REGISTRY_SERVER="{{.registry.server}}"
export REGISTRY_USERNAME="{{.registry.username}}"
export REGISTRY_PASSWORD="{{.registry.password}}"
export MYSQL_SERVER="{{.mysql.server}}"
export MYSQL_ROOT_PASSWORD="{{.mysql.rootPassword}}"
export MYSQL_USERNAME="{{.mysql.username}}"
export MYSQL_PASSWORD="{{.mysql.password}}"
export MYSQL_DATABASE="ancient"
export ELASTICSEARCH_SERVER="elasticsearch-master:9200"
export ELASTICSEARCH_INDEX="shici"
export REDIS_ADDRESS="{{.redis.addr}}"
export REDIS_PASSWORD="{{.redis.password}}"
export PULL_SECRETS="hatlonely-pull-secrets"
export IMAGE_REPOSITORY="hatlonely/rpc-ancient"
export IMAGE_TAG="$(cd .. && git describe --tags | awk '{print(substr($0,2,length($0)))}')"
export REPLICA_COUNT=2
export INGRESS_HOST="k8s.ancient.hatlonely.com"
export INGRESS_SECRET="k8s-secret"
export K8S_CONTEXT="homek8s"