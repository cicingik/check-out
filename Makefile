SHELL := /bin/bash
DOCKER_BUILDKIT = 1
BUILD_TIME = $(shell date)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG = $(shell git describe --abbrev=0 2> /dev/null || echo "no tag")
COMMIT = $(shell git log -1 --pretty=format:"%at-%h")
COMMIT_MSG = $(shell git log -1 --pretty=format:"%s")
DOCKER_IMAGE_TAG ?= "asia.gcr.io/cicingik/check-out:local-latest"

APP_ENTRYPOINT ?= ./cmd/check-out

.PHONY: run
run: ## execute go run main.go
	@go run cmd/main.go


.PHONY: lint
lint:  ## Lint this codebase
	@go mod tidy
	@golint .
	@gofmt -e -s -w .
	@goimports -v -w .

.PHONY: app-image
app-image:  ## Create a docker image
	@echo "Building docker image with tag: ${DOCKER_IMAGE_TAG}"
	@docker build \
		--rm \
 		--compress \
		--build-arg "COMMIT_MSG=${COMMIT_MSG}" \
 		--build-arg "COMMIT=${COMMIT}" \
 		--build-arg "GIT_TAG=${GIT_TAG}" \
 		--build-arg "BUILD_TAG=${BUILD_TAG}" \
		--build-arg "APP_ENTRYPOINT=${APP_ENTRYPOINT}" \
 		-t ${DOCKER_IMAGE_TAG} \
 		-f Dockerfile .


.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo  "[!] Available Command: "
	@echo  "-----------------------"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' | sort