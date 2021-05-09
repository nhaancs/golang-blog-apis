include .env

rundb:
	@docker run -d --name mysql --privileged=true -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	bitnami/mysql:8.0
buildmigrator:
	@docker build -t migrator ./migrator
startdb:
	@docker start mysql
migrateup:
	@docker run --network host migrator -path="/migrations/" -database "mysql://${DSN}" up
start:
	@go run .
fmt:
	@go fmt ./...

.PHONY: rundb startdb migrateup buildmigrator start fmt
