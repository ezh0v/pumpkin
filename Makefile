ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

DATABASE_CONNECT ?= host=$(POSTGRES_HOST) port=5432 user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable TimeZone=$(TIMEZONE)

##@ Migrations
migrate-new: ## Create a new migration. Example: make migrate-new NAME=
	goose create -dir migrations $(NAME) sql

migrate-up: ## Apply all migrations
	goose postgres "$(DATABASE_CONNECT)" up -dir migrations

migrate-down: ## Roll back the last migration
	goose postgres "$(DATABASE_CONNECT)" down -dir migrations

##@ Application
run: ## Runs the application locally using the current environment configuration
	APP_VERSION=dev \
	DATABASE_CONNECT="$(DATABASE_CONNECT)" \
	go run main.go