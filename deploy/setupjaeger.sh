#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Setting up jaeger..."

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  "docker rm -f jaeger && docker run -d --name jaeger \
    --network ${DOCKER_NETWORK} \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -e VIRTUAL_HOST=${TRACING_VIRTUAL_HOST} \
    -e VIRTUAL_PORT=${TRACING_VIRTUAL_PORT} \
    -e LETSENCRYPT_HOST=${TRACING_LETSENCRYPT_HOST} \
    -e LETSENCRYPT_EMAIL=${TRACING_LETSENCRYPT_EMAIL} \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 14250:14250 \
    -p 9411:9411 \
    jaegertracing/all-in-one:1.22"

echo "Done"
