PATH := ./migrations/
HOST := clickhouse
PORT := 9000
DATABSE := default
DATABASE_URL := 'clickhouse://${HOST}:${PORT}/${DATABSE}'
EXT := sql
MIGRATE := /go/src/github.com/golang-migrate/migrate/bin/migrate

# コマンド
## make create NAME=create_table_name
.PHONY: create
create:
	${MIGRATE} create -ext ${EXT} -dir ${PATH} ${NAME}

.PHONY: up
up:
	${MIGRATE} -database ${DATABASE_URL} -path ${PATH} up

.PHONY: down
down:
	${MIGRATE} -database ${DATABASE_URL} -path ${PATH} down
