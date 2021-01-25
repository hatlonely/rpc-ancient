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
      "http": {
        "port": 80
      },
      "grpc": {
        "port": 6080
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
              "filename": "log/${NAME}.rpc",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
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
              "filename": "log/${NAME}.rpc",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
              }
            }
          }]
        }
      }
    }
