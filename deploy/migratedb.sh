#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Docker building migrator image..."
docker build -t migrator -f Dockerfile ./migrator
echo "Docker saving migrator image..."
docker save -o migrator.tar migrator

echo "Deploying migrator image to server..."
scp -o StrictHostKeyChecking=no ./migrator.tar ${DEPLOY_CONNECT}:~
ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/prod/migratedb.sh

echo "Cleaning..."
rm -f ./migrator.tar
echo "Done"