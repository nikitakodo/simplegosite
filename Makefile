.PHONY: install_dev
install_dev:
		./environment/init.d/install_dev.sh

.PHONY: start
start:
		docker-compose up --build -d

.PHONY: stop
stop:
		docker-compose stop

.PHONY: prune
prune:
		docker-compose down -v

.PHONY: migrate_up
migrate_up:
		docker-compose exec app ./migration -mode=up

.PHONY: migrate_down
migrate_down:
		docker-compose exec app ./migration -mode=down

.PHONY: clean_cache
clean_cache:
		docker exec -it redis redis-cli FLUSHALL
.PHONY: new_migration
new_migration:
		./scripts/create_migration.sh $(filter-out $@,$(MAKECMDGOALS))
%:
		@:

.PHONY: build
build:
		make start && make migrate_up
