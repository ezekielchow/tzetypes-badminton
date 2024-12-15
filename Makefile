.PHONY: openapi test sqlc-gen create-migration gen-mocks pre-commit openapi-js lint web-build

help: ## Show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

openapi: ## Generate all openapis
	docker build -t oapi-codegen-docker -f ./docker/oapi-codegen/Dockerfile . 
	@./scripts/openapi-http.sh private internal/common/oapiprivate oapiprivate
	@./scripts/openapi-http.sh public internal/common/oapipublic oapipublic

openapi-js:  ## Generate openapis for frontend
	@./scripts/openapi-js.sh private
	@./scripts/openapi-js.sh public

test:  ## Run tests
	@./scripts/test.sh common .test.env
	@./scripts/test.sh users .test.env
	@./scripts/test.sh sessions .test.env
	@./scripts/test.sh players .test.env
	@./scripts/test.sh games .test.env

sqlc-gen: ## Generate golang code from sqlc queries
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate --file internal/database/sqlc.yml

create-migration: ## Create migration file; Usage: make create-migration table_name=your_table_name
	docker run --rm --user $(id -u):$(id -g) -v $(PWD)/internal/database/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq ${table_name}

gen-mocks:  ## Generate mocks to help in testing
	docker run --rm -v ./internal:/src -w /src vektra/mockery --all

web-build: ## Build Frontend
	cd ./web && bun run build

pre-commit: openapi openapi-js sqlc-gen gen-mocks lint test web-build ## Make sure all code is ok before commiting

lint: ## Lint golang
	docker run --rm -v ./internal:/app -v ~/.cache/golangci-lint/v1.61.0:/root/.cache -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v ./...
