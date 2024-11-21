MIGRATE=./migrations/migrate
DB_URL=mysql://root:@tcp(localhost:3306)/my_golang_db
MIGRATIONS_PATH=./migrations/sql

migrate-up:
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

migrate-down:
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATIONS_PATH) down 1

migrate-create:
	@read -p "Enter migration name: " name; \
	mkdir -p $(MIGRATIONS_PATH) && \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_PATH) $$name

.PHONY: seed
seed:
	@echo "Running seeders..."
	go run main.go seed

.PHONY: migration
migration:
	@echo "Running migration..."
	go run main.go migration