include .env

rundb:
	@docker run -d --name mysql --privileged=true -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	-v ~/mysql:/bitnami \
	bitnami/mysql:8.0
buildmigrator:
	@docker build -t migrator ./migrator
startdb:
	@docker start mysql
migrateup:
	@docker run --network host migrator -path="/migrations/" -database "mysql://${DSN}" up
start:
	@PORT="${PORT}" \
	GIN_MODE="${GIN_MODE}" \
	DSN="${DSN}" \
	AUTH_SECRET="${AUTH_SECRET}" \
	S3_BUCKET_NAME="${S3_BUCKET_NAME}" \
	S3_REGION="${S3_REGION}" \
	S3_API_KEY="${S3_API_KEY}" \
	S3_SECRET_KEY="${S3_SECRET_KEY}" \
	S3_DOMAIN="${S3_DOMAIN}" \
	go run .
fmt:
	@go fmt ./...
deploy:
	@./deploy.sh

.PHONY: rundb startdb migrateup buildmigrator start fmt
