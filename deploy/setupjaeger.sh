#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Setting up jaeger"

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  "docker rm -f jaeger && docker run -d --name jaeger \
    --network ${DOCKER_NETWORK} \
    -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -e VIRTUAL_HOST="jaeger.nhannguyen.codes" \
    -e VIRTUAL_PORT=16686 \
    -e LETSENCRYPT_HOST="jaeger.nhannguyen.codes" \
    -e LETSENCRYPT_EMAIL="nhanpublic@gmail.com" \
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
