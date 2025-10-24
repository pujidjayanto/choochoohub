.PHONY: migrate rollback create-migration seed create-seed run mocks

# Run migrations up
migrate:
	go run ./cmd/migration -migrate -- up

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

# Generate mocks for all interfaces in key directories
mocks:
	@echo "Cleaning old mocks..."
	@rm -rf ./mocks

	@echo "Generating mocks into ./mocks ..."
	@mkdir -p ./mocks

	# ----- API -----
	@mockgen -source=./api/ping.go -package=mocks -destination=./mocks/mock_ping.go
	@mockgen -source=./api/signin.go -package=mocks -destination=./mocks/mock_signin.go
	@mockgen -source=./api/signup.go -package=mocks -destination=./mocks/mock_signup.go

	# ----- Repository -----
	@mockgen -source=./repository/user.go -package=mocks -destination=./mocks/mock_user_repository.go
	@mockgen -source=./repository/user_test.go -package=mocks -destination=./mocks/mock_user_test.go

	# ----- Usecase -----
	@mockgen -source=./usecase/signin.go -package=mocks -destination=./mocks/mock_signin_usecase.go
	@mockgen -source=./usecase/signup.go -package=mocks -destination=./mocks/mock_signup_usecase.go

	# ----- Pkg -----
	@mockgen -source=./pkg/db/gorm.go -package=mocks -destination=./mocks/mock_db_gorm.go
	@mockgen -source=./pkg/envloader/envloader.go -package=mocks -destination=./mocks/mock_envloader.go
	@mockgen -source=./pkg/logger/logrus.go -package=mocks -destination=./mocks/mock_logger.go
	@mockgen -source=./pkg/pwd/pwd.go -package=mocks -destination=./mocks/mock_pwd.go

	@echo "All mocks generated in ./mocks"
