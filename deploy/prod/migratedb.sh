#!/usr/bin/env bash

docker load -i migrator.tar
docker rm -f migrator

docker run \
    --network ${DOCKER_NETWORK} \
    --name migrator \
    migrator \
    -path="/migrations/" \
    -database "mysql://$(echo "${DSN}" | base64 --decode)" \
    up