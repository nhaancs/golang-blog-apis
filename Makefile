include .env

createnetwork: 
	@docker network create ${DOCKER_NETWORK}
rundb:
	@docker run -d \
	--network ${DOCKER_NETWORK} \
	--name mysql \
	--privileged=true \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	bitnami/mysql:8.0
startdb:
	@docker start mysql
migrateup:
	@docker build -t migrator ./migrator && \
	docker rm -f migrator && \
	docker run \
	--name migrator \
	--network ${DOCKER_NETWORK} \
	migrator \
	-path="/migrations/" \
	-database "mysql://${DSN}" \
	up
start:
	@PORT="${PORT}" \
	GIN_MODE="${GIN_MODE}" \
	DSN="${DSN_LOCAL}" \
	AUTH_SECRET="${AUTH_SECRET}" \
	S3_BUCKET_NAME="${S3_BUCKET_NAME}" \
	S3_REGION="${S3_REGION}" \
	S3_API_KEY="${S3_API_KEY}" \
	S3_SECRET_KEY="${S3_SECRET_KEY}" \
	S3_DOMAIN="${S3_DOMAIN}" \
	TRACING_AGENT_ENPOINT=${TRACING_AGENT_ENPOINT} \
	TRACING_PROBABILITY_SAMPLER=${TRACING_PROBABILITY_SAMPLER} \
	TRACING_APP_NAME=${TRACING_APP_NAME} \
	go run .
fmt:
	@go fmt ./...

setpermissions:
	@chmod +x ./deploy/deploy.sh ./deploy/migratedb.sh ./deploy/setupserver.sh
setupserver:
	@./deploy/setupserver.sh
setupjaeger:
	@./deploy/setupjaeger.sh
setupnetdata:
	@./deploy/setupnetdata.sh
loadtest:
	@# https://medium.com/@felipedutratine/intelligent-benchmark-with-wrk-163986c1587f
	@wrk -t6 -c200 -d30s https://rest.nhannguyen.codes/v1/categories
migratedb:
	@./deploy/migratedb.sh
deploy:
	@./deploy/deploy.sh

.PHONY: rundb startdb migrateup start fmt deploy migratedb setupserver setpermissions setupjaeger setupnetdata
