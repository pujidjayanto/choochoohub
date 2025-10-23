.PHONY: migrate rollback create-migration seed create-seed run

# Run migrations up
migrate:
	go run ./cmd/migration -migrate up

# Rollback migrations (last batch)
rollback:
	go run ./cmd/migration -migrate down

# Create a new migration
create-migration:
	@echo "Usage: make create-migration NAME=your_migration_name"
	@go run ./cmd/migration -create-migration $(NAME)

# Run seeds
seed:
	go run ./cmd/migration -seed

# Create a new seed
create-seed:
	@echo "Usage: make create-seed NAME=your_seed_name"
	@go run ./cmd/migration -create-seed $(NAME)

# Run the user-api
run:
	go run ./cmd/user-api
