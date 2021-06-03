#!/usr/bin/env bash

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d \
  --name ${APP_NAME} \ 
  --network ${DOCKER_NETWORK} \
  -p ${EXPORTED_PORT}:${PORT} \
  -e PORT="${PORT}" \
  -e GIN_MODE="${GIN_MODE}" \
  -e VIRTUAL_HOST="${VIRTUAL_HOST}" \
  -e LETSENCRYPT_HOST="${LETSENCRYPT_HOST}" \
  -e LETSENCRYPT_EMAIL="${LETSENCRYPT_EMAIL}" \
  -e DSN="${DSN}" \
  -e AUTH_SECRET="${AUTH_SECRET}" \
  -e S3_BUCKET_NAME="${S3_BUCKET_NAME}" \
  -e S3_REGION="${S3_REGION}" \
  -e S3_API_KEY="${S3_API_KEY}" \
  -e S3_SECRET_KEY="${S3_SECRET_KEY}" \
  -e S3_DOMAIN="${S3_DOMAIN}" \
  ${APP_NAME}