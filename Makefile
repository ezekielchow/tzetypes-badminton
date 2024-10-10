.PHONY: openapi test sqlc-gen create-migration gen-mocks pre-commit openapi-js

help: ## Show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

openapi: ## Generate all openapis
	@./scripts/openapi-http.sh private internal/common/oapiprivate oapiprivate
	@./scripts/openapi-http.sh public internal/common/oapipublic oapipublic

openapi-js:  ## Generate openapis for frontend
	@./scripts/openapi-js.sh private
	@./scripts/openapi-js.sh public

test:  ## Run tests
	@./scripts/test.sh common .test.env
	@./scripts/test.sh users .test.env
	@./scripts/test.sh sessions .test.env

sqlc-gen: ## Generate golang code from sqlc queries
	sqlc generate --file internal/users/store/sqlc.yml
	sqlc generate --file internal/sessions/store/sqlc.yml

create-migration: ## Create migration file; Usage: make create-migration table_name=your_table_name
	migrate create -ext sql -dir ./internal/migrations -seq $(table_name)

gen-mocks:  ## Generate mocks to help in testing
	docker run -v ./internal:/src -w /src vektra/mockery --all

pre-commit: openapi sqlc-gen gen-mocks test ## Make sure all code is ok before commiting
