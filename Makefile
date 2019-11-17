.PHONY: start
start:
		docker-compose up --build -d

.PHONY: stop
stop:
		docker-compose stop

.PHONY: migrate_up
migrate_up:
		docker-compose exec app ./migration -mode=up

.PHONY: migrate_down
migrate_down:
		docker-compose exec app ./migration -mode=down

.PHONY: clean_cache
clean_cache:
		docker-compose exec redis sh
.PHONY: new_migration
new_migration:
		./scripts/create_migration.sh $(filter-out $@,$(MAKECMDGOALS))
%:
		@:

.PHONY: build
build:
		make start && make migrate_up

.PHONY: install_migrate
install_migrate:
		./scripts/install_migrate.sh

.DEFAULT_GOAL := build