
.PHONY: run
run: ## execute go run main.go
	@go run cmd/main.go


.PHONY: lint
lint:  ## Lint this codebase
	@go mod tidy
	@golint .
	@gofmt -e -s -w .
	@goimports -v -w .


.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo  "[!] Available Command: "
	@echo  "-----------------------"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' | sort