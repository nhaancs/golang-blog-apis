#!/usr/bin/env bash

echo "Installing Docker..."
curl -fsSL https://get.docker.com -o get-docker.sh && chmod +x ./get-docker.sh && ./get-docker.sh

echo "Creating Docker network..."
docker network create ${DOCKER_NETWORK}

echo "Starting database container..."
docker rm -f mysql && docker run -d \
  --user root \
  -p 3306:3306 \
  --network ${DOCKER_NETWORK} \
  --name mysql \
  --privileged=true \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
  -v ~/mysql_data:/bitnami/mysql/data \
	bitnami/mysql:8.0

echo "Starting nginx proxy container..."
docker rm -f nginx-proxy && docker run -d \
  -p 80:80 -p 443:443 \
  --network ${DOCKER_NETWORK} \
  --name nginx-proxy \
  --label nginx_proxy \
  -e ENABLE_IPV6=true \
  --privileged=true \
  -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
  -v ~/nginx-certs:/etc/nginx/certs:ro \
  -v ~/nginx-conf:/etc/nginx/conf.d \
  -v ~/nginx-logs:/var/log/nginx \
  -v /usr/share/nginx/html \
  -v /var/run/docker.sock:/tmp/docker.sock:ro \
  jwilder/nginx-proxy

echo "Starting Letsencrypt container..."
docker rm -f letsencrypt && docker run -d \
  --name letsencrypt \
  --network ${DOCKER_NETWORK} \
  --privileged=true \
  -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
  -v ~/nginx-certs:/etc/nginx/certs:rw \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --volumes-from nginx-proxy \
  jrcs/letsencrypt-nginx-proxy-companion
