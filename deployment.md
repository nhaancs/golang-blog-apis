# Deployment

## Server setup (first time only)
- Setup ssh connection from client to server
- [Install docker](https://docs.docker.com/engine/install/ubuntu/)
- Create Docker network
    ```bash
    docker network create ${DOCKER_NETWORK}
    ```
- Run database container
    ```bash
    docker run -d \
    -p 3306:3306 \
    --network ${DOCKER_NETWORK} \
    --name mysql \
    --privileged=true \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	-v ~/mysql:/bitnami \
	bitnami/mysql:8.0
    ``
- Run nginx proxy container
    ```bash
    docker run -d \
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
    ```
- Run Letsencrypt container
    ```bash
    docker run -d \
    --network ${DOCKER_NETWORK} \
    --privileged=true \
    -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
    -v ~/nginx-certs:/etc/nginx/certs:rw \
    -v /var/run/docker.sock:/var/run/docker.sock:ro \
    --volumes-from nginx-proxy \
    jrcs/letsencrypt-nginx-proxy-companion
    ```

## Client setup
- Provide permission to run `deploy.sh`
    ```bash
    chmod +x deploy.sh
    ```
- Update environment variables in `.env` file

## Deploy
