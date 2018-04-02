DC_FILE := -f env/docker-compose.yml


.PHONY: env
env:
	cp -n env/.example.env env/.env || true # for containers
	cp -n env/.env .env             || true # for docker compose file, https://docs.docker.com/compose/env-file/

.PHONY: reset-env
reset-env:
	rm -f env/.env env/cross-origin/.env .env || true


.PHONY: up
up: env
	docker-compose $(DC_FILE) up -d
	docker-compose $(DC_FILE) rm -f

.PHONY: fresh-up
fresh-up: env
	docker-compose $(DC_FILE) up --build --force-recreate -d
	docker-compose $(DC_FILE) rm -f

.PHONY: down
down: env
	docker-compose $(DC_FILE) down

.PHONY: clean-down
clean-down: env
	docker-compose $(DC_FILE) down --volumes --rmi local

.PHONY: clear
clear: env
	docker-compose $(DC_FILE) rm -f

.PHONY: status
status: env
	docker-compose $(DC_FILE) ps


.PHONY: up-db
up-db:
	docker-compose $(DC_FILE) up -d db

.PHONY: start-db
start-db: env
	docker-compose $(DC_FILE) start db

.PHONY: stop-db
stop-db: env
	docker-compose $(DC_FILE) stop db

.PHONY: log-db
log-db: env
	docker-compose $(DC_FILE) logs -f db

.PHONY: psql
psql: env
	docker-compose $(DC_FILE) exec db /bin/sh -c 'su - postgres -c psql'

.PHONY: backup
backup: env
	docker-compose $(DC_FILE) exec db \
	  /bin/sh -c 'su - postgres -c "pg_dump --clean $${POSTGRES_DB}"' > ./env/backup.sql

.PHONY: restore
restore:
	cat ./env/backup.sql | docker exec -i $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1) \
	  /bin/sh -c 'cat > /tmp/backup.sql && su - postgres -c "psql $${POSTGRES_DB} < /tmp/backup.sql"'


.PHONY: up-migration
up-migration:
	docker-compose $(DC_FILE) up --build -d migration

.PHONY: start-migration
start-migration: env
	docker-compose $(DC_FILE) start migration

.PHONY: log-migration
log-migration: env
	docker-compose $(DC_FILE) logs -f migration


.PHONY: up-service
up-service:
	docker-compose $(DC_FILE) up --build -d service

.PHONY: start-service
start-service: env
	docker-compose $(DC_FILE) start service

.PHONY: stop-service
stop-service: env
	docker-compose $(DC_FILE) stop service

.PHONY: log-service
log-service: env
	docker-compose $(DC_FILE) logs -f service


.PHONY: up-server
up-server:
	docker-compose $(DC_FILE) up -d server

.PHONY: start-server
start-server: env
	docker-compose $(DC_FILE) start server

.PHONY: stop-server
stop-server: env
	docker-compose $(DC_FILE) stop server

.PHONY: log-server
log-server: env
	docker-compose $(DC_FILE) logs -f server

# ~~~

.PHONY: cross-origin-up
cross-origin-up:
	cp -n env/cross-origin/.example.env env/cross-origin/.env || true
	cp -n env/cross-origin/.env .env                          || true
	docker-compose -f env/cross-origin/docker-compose.yml up -d

.PHONY: cross-origin-status
cross-origin-status:
	cp -n env/cross-origin/.example.env env/cross-origin/.env || true
	cp -n env/cross-origin/.env .env                          || true
	docker-compose -f env/cross-origin/docker-compose.yml ps

.PHONY: cross-origin-down
cross-origin-down:
	cp -n env/cross-origin/.example.env env/cross-origin/.env || true
	cp -n env/cross-origin/.env .env                          || true
	docker-compose -f env/cross-origin/docker-compose.yml down --volumes --rmi local
