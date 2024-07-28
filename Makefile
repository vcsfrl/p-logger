# HELP
.PHONY: help

help: ## Usage: make <option>
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies.
	go mod vendor
	go mod tidy

run-example: ## Run example.
	go run example/main.go

test: ## Run all tests.
	go test ./... -race -cpu 24 -cover -coverprofile=var/tests/coverage.out;

test-cover: test ## Test coverage.
	go tool cover -func=var/tests/coverage.out


dc-install: ## Install with docker compose.
	if [ ! -f .env ]; then echo "CONTAINER_EXEC_USER_ID=`id -u`" >> .env; echo "CONTAINER_USERNAME=${USER}" >> .env; fi
	docker compose build;
	docker compose run p_logger make install

dc-example: ## Run example with docker compose.
	docker compose run p_logger make run-example

dc-test: ## Run all tests with docker compose.
	docker compose run p_logger make test

dc-test-cover: ## Test coverage with docker compose.
	docker compose run p_logger make test-cover

