#!/bin/sh

apk --no-cache add postgresql-client;

while ! pg_isready -h $DB_HOST -p $DB_PORT; do sleep 1; done;

# 1. Если на дев-машине, можешь просто закинуть бинарник в папку migrator.
#/migrator/migrate.linux-amd64 -path /migrations -database postgres://$DB_USER:$DB_PASSWORD@postgres:$DB_PORT/$DB_NAME?sslmode=disable up

# 2. если хочешь, загрузи новую версию мигратора.
apk --no-cache add curl;
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.6.2/migrate.linux-amd64.tar.gz | tar xvz;
./migrate.linux-amd64 -path /migrations -database postgres://$DB_USER:$DB_PASSWORD@postgres:$DB_PORT/$DB_NAME?sslmode=disable up
