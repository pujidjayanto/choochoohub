.PHONY: migrate rollback create-migration seed create-seed

# Run migrations up
migrate:
	go run ./cmd/migration/main.go -migrate up

# Rollback migrations (last batch)
rollback:
	go run ./cmd/migration/main.go -migrate down

# Create a new migration
create-migration:
	@echo "Usage: make create-migration NAME=your_migration_name"
	@go run ./cmd/migration/main.go -create-migration $(NAME)

# Run seeds
seed:
	go run ./cmd/migration/main.go -seed

# Create a new seed
create-seed:
	@echo "Usage: make create-seed NAME=your_seed_name"
	@go run ./cmd/migration/main.go -create-seed $(NAME)
