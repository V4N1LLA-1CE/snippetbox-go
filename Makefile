# load env files
include .env
export

# create the container
postgres:
	docker run --name ${POSTGRES_CONTAINER_NAME} \
	-p 5432:5432 \
	-e POSTGRES_USER=${POSTGRES_USER} \
	-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
	-d postgres:latest

# create db inside container
createdb:
	docker exec -it ${POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB_NAME}

# drop db inside container
dropdb:
	docker exec -it ${POSTGRES_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB_NAME}

# migrate db up
migrateup:
	migrate -path db/migrations -database "${POSTGRES_URL}" -verbose up

# migrate db down
migratedown:
	migrate -path db/migrations -database "${POSTGRES_URL}" -verbose down

# run app with reloading
watch:
	air

.PHONY: watch postgres createdb dropdb migrateup migratedown
