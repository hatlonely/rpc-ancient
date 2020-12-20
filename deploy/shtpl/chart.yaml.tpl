namespace: "${NAMESPACE}"
name: "${NAME}"
replicaCount: "${REPLICA_COUNT}"

image:
  repository: "${REGISTRY_SERVER}/${IMAGE_REPOSITORY}"
  tag: "${IMAGE_TAG}"
  pullPolicy: Always

imagePULL_SECRETS:
  name: "${PULL_SECRETS}"

ingress:
  host: "${INGRESS_HOST}"
  secretNAME: "${INGRESS_SECRET}"

config: |
  {
    "http": {
      "port": 80
    },
    "grpc": {
      "port": 6080
    },
    "mysql": {
      "username": "${MYSQL_USERNAME}",
      "password": "${MYSQL_PASSWORD}",
      "database": "${MYSQL_DATABASE}",
      "host": "${MYSQL_SERVER}",
      "port": 3306,
      "connMaxLifeTime": "60s",
      "maxIdleConns": 10,
      "maxOpenConns": 20
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
            "elasticSearch": {
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