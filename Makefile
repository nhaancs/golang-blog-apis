include .env

run-mysql:
	docker run -d --name mysql --privileged=true -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
	-e MYSQL_DATABASE=${MYSQL_DATABASE} \
	bitnami/mysql:8.0
start-mysql:
	docker start mysql
stop-mysql:
	docker stop mysql
start-server:
	go run .