.DEFAULT_GOAL := help

.PHONY: help
help: ## List user-facing Make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: codegen
codegen: ## Generate provider code
	go generate ./...

.PHONY: build
build: codegen ## Build the provider binary
	go install .

.PHONY: test
test: ## Run tests
	go test ./...
