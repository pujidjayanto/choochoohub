.PHONY: run lint

run:
	go run ./cmd/api-gateway

lint:
	@echo "Running golint..."
	golangci-lint run ./...
