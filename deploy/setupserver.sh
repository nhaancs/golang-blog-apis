#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Setting up server..."

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  DOCKER_NETWORK=${DOCKER_NETWORK} \
  DB_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
  DB_USER=${DB_USER} \
  DB_PASSWORD=${DB_PASSWORD} \
  DB_NAME=${DB_NAME} \
  'bash -s' < ./deploy/prod/setupserver.sh

echo "Done"
