namespace: "${NAMESPACE}"
name: "${NAME}"
replicaCount: "${REPLICA_COUNT}"

image:
  repository: "${REGISTRY_SERVER}/${REGISTRY_NAMESPACE}/${IMAGE_REPOSITORY}"
  tag: "${IMAGE_TAG}"
  pullPolicy: Always
  pullSecret: "${IMAGE_PULL_SECRET}"

ingress:
  enable: true
  host: "${INGRESS_HOST}"
  secretName: "${INGRESS_SECRET}"

config:
  app: |
    {
      "grpcGateway": {
        "httpPort": 80,
        "grpcPort": 6080,
        "enableTrace": true,
        "enableMetric": true,
        "enablePprof": true,
        "validators": ["Default"],
        "usePascalNameLogKey": false,
        "usePascalNameErrKey": false,
        "marshalUseProtoNames": true,
        "marshalEmitUnpopulated": false,
        "unmarshalDiscardUnknown": true,
        "jaeger": {
          "serviceName": "rpc-ancient",
          "sampler": {
            "type": "const",
            "param": 1
          },
          "reporter": {
            "logSpans": false
          }
        },
        "rateLimiterHeader": "x-user-id",
        "rateLimiter": {
          "type": "RedisRateLimiterInstance"
        },
        "parallelControllerHeader": "x-user-id",
        "parallelController": {
          "type": "RedisTimedParallelControllerInstance"
        }
      },
      "parallelController": {
        "redis": {
          "redis": {
            "addr": "${REDIS_ADDRESS}"
          },
          "wrapper": {
            "enableTrace": true,
            "enableMetric": true
          }
        },
        "defaultMaxToken": 3,
        "maxToken": {
          "123|/api.AncientService/GetAncient": 3
        },
        "interval": "3s",
        "expiration": "10s"
      },
      "rateLimiter": {
        "redis": {
          "redis": {
            "addr": "${REDIS_ADDRESS}"
          },
          "wrapper": {
            "enableTrace": true,
            "enableMetric": true
          }
        },
        "qps": {
          "123|/api.AncientService/GetAncient": 4
        }
      },
      "mysql": {
        "gorm": {
          "username": "${MYSQL_USERNAME}",
          "password": "${MYSQL_PASSWORD}",
          "database": "${MYSQL_DATABASE}",
          "host": "${MYSQL_SERVER}",
          "port": 3306,
          "connMaxLifeTime": "60s",
          "maxIdleConns": 10,
          "maxOpenConns": 20
        },
        "wrapper": {
          "enableTrace": true
        },
        "retry": {
          "attempts": 3,
          "lastErrorOnly": true,
          "delayType": "BackOff"
        }
      },
      "elasticsearch": {
        "uri": "http://${ELASTICSEARCH_SERVER}"
      },
      "service": {
        "elasticsearchIndex": "shici"
      },
      "logger": {
        "grpc": {
          "level": "Info",
          "flatMap": true,
          "writers": [{
            "type": "RotateFile",
            "rotateFileWriter": {
              "filename": "log/app.rpc",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
                "options": {
                  "flatMap": true
                }
              }
            }
          }, {
            "type": "ElasticSearch",
            "elasticSearchWriter": {
              "index": "grpc",
              "idField": "requestID",
              "timeout": "200ms",
              "msgChanLen": 200,
              "workerNum": 2,
              "es": {
                "uri": "http://${ELASTICSEARCH_SERVER}"
              }
            }
          }]
        },
        "info": {
          "level": "Info",
          "writers": [{
            "type": "RotateFile",
            "rotateFileWriter": {
              "filename": "log/app.log",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
              }
            }
          }]
        }
      }
    }
