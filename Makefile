SHELL := /bin/bash
DOCKER_BUILDKIT = 1
BUILD_TIME = $(shell date)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG = $(shell git describe --abbrev=0 2> /dev/null || echo "no tag")
COMMIT = $(shell git log -1 --pretty=format:"%at-%h")
COMMIT_MSG = $(shell git log -1 --pretty=format:"%s")
DOCKER_IMAGE_TAG ?= "asia.gcr.io/cicingik/check-out:local-latest"

APP_PID      = /tmp/checkout.pid
APP_NAME        = ./checkout
APP_ENTRYPOINT ?= ./cmd/checkout
BUILD_TAG      ?= development

.PHONY: restart
serve: restart  ## Run application and automaticaly restart on source code change
	@fswatch --event=Updated -or -e ".*" -i ".*/[^.]*\\.go$$" ./internal ./app ./api ./config  ./domain ./pkg ./usecase ./repository | xargs -n1 -I{}  make restart || make kill

kill:
	@echo "killing old checkout instance"
	@kill `cat $(APP_PID)` >> /dev/null 2>&1 || true


restart: kill compile_dev
	@"$(APP_NAME)-$(BUILD_TAG)" & echo $$! > $(APP_PID)


.PHONY: run
run: ## execute go run main.go
	@go run cmd/checkout/main.go


.PHONY: lint
lint:  ## Lint this codebase
	@go mod tidy
	@golint .
	@gofmt -e -s -w .
	@goimports -v -w .


compile_dev:  ## Compile dev version of application
	@echo "Compiling application with tag: ${BUILD_TAG}..."
	CGO_ENABLED=1 go build \
		 -i -v -race  \
		-ldflags="\
			-X \"github.com/cicingik/check-out/config.BuildTime=${BUILD_TIME}\" \
			-X \"github.com/cicingik/check-out/config.CommitMsg=${COMMIT_MSG}\" \
			-X \"github.com/cicingik/check-out/config.CommitHash=${COMMIT}\" \
			-X \"github.com/cicingik/check-out/config.AppVersion=${GIT_TAG}\" \
			-X \"github.com/cicingik/check-out/config.ReleaseVersion=${BUILD_TAG}\"" \
		-tags ${BUILD_TAG} \
		-o $(APP_NAME)-$(BUILD_TAG) \
		${APP_ENTRYPOINT}


compile:  ## Build binary version of application
	@echo "Compiling application with tag: ${BUILD_TAG}..."
	CGO_ENABLED=1 go build \
		 -i -v -race -a  \
		-ldflags="\
			-X \"github.com/cicingik/check-out/config.BuildTime=${BUILD_TIME}\" \
			-X \"github.com/cicingik/check-out/config.CommitMsg=${COMMIT_MSG}\" \
			-X \"github.com/cicingik/check-out/config.CommitHash=${COMMIT}\" \
			-X \"github.com/cicingik/check-out/config.AppVersion=${GIT_TAG}\" \
			-X \"github.com/cicingik/check-out/config.ReleaseVersion=${BUILD_TAG}\"" \
		-tags production \
		-o $(APP_NAME)-production \
		${APP_ENTRYPOINT}


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