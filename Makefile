
MIGRATE=migrate
MIGRATIONS_DIR=database/migrations
DB_DSN=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

# install dev tools
tool:
	brew install golang-migrate@v4.18.3
# Load env vars

## Run migrations up
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" up

## Rollback one migration
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" down 1

run:
	go run cmd/main.go