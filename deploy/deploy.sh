#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Docker building..."
docker build -t ${APP_NAME} -f ./Dockerfile .
echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}

echo "Deploying..."
scp -o StrictHostKeyChecking=no ./${APP_NAME}.tar ${DEPLOY_CONNECT}:~

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  APP_NAME=${APP_NAME} \
  DOCKER_NETWORK=${DOCKER_NETWORK} \
  EXPORTED_PORT=${EXPORTED_PORT} \
  PORT=${PORT} \
  GIN_MODE=${GIN_MODE} \
  VIRTUAL_HOST=${VIRTUAL_HOST} \
  LETSENCRYPT_HOST=${LETSENCRYPT_HOST} \
  LETSENCRYPT_EMAIL=${LETSENCRYPT_EMAIL} \
  DSN=$(echo -ne $DSN | base64 -w 0) \
  AUTH_SECRET=${AUTH_SECRET} \
  S3_BUCKET_NAME=${S3_BUCKET_NAME} \
  S3_REGION=${S3_REGION} \
  S3_API_KEY=${S3_API_KEY} \
  S3_SECRET_KEY=${S3_SECRET_KEY} \
  S3_DOMAIN=${S3_DOMAIN} \
  TRACING_AGENT_ENPOINT=${TRACING_AGENT_ENPOINT} \
  TRACING_PROBABILITY_SAMPLER=${TRACING_PROBABILITY_SAMPLER} \
  TRACING_APP_NAME=${TRACING_APP_NAME} \
  'bash -s' < ./deploy/prod/start.sh

echo "Cleaning..."
rm -f ./${APP_NAME}.tar
echo "Done"