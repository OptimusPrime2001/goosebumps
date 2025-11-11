include .env
export 

MIGRATE_PATH=internal/db/migrations
# CONN_STR=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
CONN_STR=$(DATABASE_URL)
ENV_FILE = .env
COMPOSE = docker-compose.yml
sqlc:
	sqlc generate

# Import database 
importdb:
	docker exec -i postgres_db psql -U $(DB_USER) -d $(DB_NAME) < user_manage.sql

# Export database 
exportdb:
	docker exec -i postgres_db pg_dump -U $(DB_USER) -d $(DB_NAME) > user_manage.sql

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATE_PATH) --database $(CONN_STR) up
# Rollback last migration
migrate-down:
	migrate -path $(MIGRATE_PATH) --database $(CONN_STR) down 1

# Force migration to a specific version
migrate-force:
	migrate -path $(MIGRATE_PATH) --database $(CONN_STR) force $(VERSION)

# Drop all migrations
migrate-drop:
	migrate -path $(MIGRATE_PATH) --database $(CONN_STR) drop

# Migrate go to a specific version
migrate-go:
	migrate -path $(MIGRATE_PATH) --database $(CONN_STR) goto $(VERSION)

make server:
	air
prod:
	docker-compose -f $(COMPOSE) down
	docker-compose -f $(COMPOSE) up -env-file $(ENV_FILE) up -d --build
logs-prod:
	docker-compose -f $(COMPOSE) logs -f --tail 100
bash:
  docker exec -it golang_api /bin/sh
.PHONY: importdb exportdb migrate-create migrate-up migrate-down migrate-force migrate-drop migrate-go sqlc server prod logs-prod bash	
