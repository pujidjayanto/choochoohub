.PHONY: migrate rollback create-migration seed create-seed run mock test

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

# Generate mocks
mock:
	@echo "Cleaning old mocks..."
	@rm -rf ./mocks
	@mkdir -p ./mocks
	@echo "Generating mocks using moq..."

	@moq -out mocks/mock_user_repository.go -pkg mocks ./repository UserRepository
	@moq -out mocks/mock_user_otp_repository.go -pkg mocks ./repository UserOtpRepository

	@moq -out mocks/mock_user_usecase.go -pkg mocks ./usecase UserUsecase
	@moq -out mocks/mock_otp_usecase.go -pkg mocks ./usecase OtpUsecase

	@moq -out mocks/mock_user_api.go -pkg mocks ./api UserApi
	@moq -out mocks/mock_otp_api.go -pkg mocks ./api OtpApi

	@moq -out mocks/mock_pkg_db_gorm.go -pkg mocks ./pkg/db/ DatabaseHandler
	@moq -out mocks/mock_pkg_eventbus.go -pkg mocks ./pkg/eventbus/ EventBus
	@moq -out mocks/mock_pkg_kafkaproducer.go -pkg mocks ./pkg/kafkaproducer/ Producer

	@echo "All mocks generated in ./mocks"


test:
	@echo "Running tests..."
	go test -v ./...

lint:
	@echo "Running golint..."
	golangci-lint run ./...
