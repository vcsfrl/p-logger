# HELP
.PHONY: help

help: ## Usage: make <option>
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies.
	go mod vendor
	go mod tidy

example: ## Run example.
	go run example/main.go

test: ## Run all tests.
	go test ./... -race -cpu 24 -cover -coverprofile=var/tests/coverage.out;

test-cover: test ## Test coverage.
	go tool cover -func=var/tests/coverage.out

