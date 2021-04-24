include .env

rundb:
	docker run -d --name mysql --privileged=true -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	bitnami/mysql:8.0
startdb:
	docker start mysql
stopdb:
	docker stop mysql
migrateup:
	migrate -database "${DB_URL}" -path "./migration/" -verbose up
start:
	go run .

.PHONY: encode rundb startdb stopdb migrateup migratedown startserver