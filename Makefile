# Define the default target
.DEFAULT_GOAL := start

# Define variables
UI_DIR := ui
GO_SERVER := cmd/server/main.go  # Replace with the entry point of your Go server

# Define targets for the Next.js app
.PHONY: install-ui
install-ui:
	@echo "Installing dependencies in $(UI_DIR)..."
	cd $(UI_DIR) && yarn install

.PHONY: start-ui
start-ui: install-ui
	@echo "Starting Next.js in development mode in $(UI_DIR)..."
	cd $(UI_DIR) && yarn run dev

.PHONY: build-ui
build-ui: install-ui
	@echo "Building the Next.js app in $(UI_DIR)..."
	cd $(UI_DIR) && yarn run build

.PHONY: start-ui-prod
start-ui-prod: build-ui
	@echo "Starting the Next.js app in production mode in $(UI_DIR)..."
	cd $(UI_DIR) && yarn start

.PHONY: lint-ui
lint-ui: install-ui
	@echo "Running linter in $(UI_DIR)..."
	cd $(UI_DIR) && yarn run lint

.PHONY: clean-ui
clean-ui:
	@echo "Cleaning up in $(UI_DIR)..."
	rm -rf $(UI_DIR)/node_modules $(UI_DIR)/.next

.PHONY: all start-db

# Default target
all: start-db

# Target to start the PostgreSQL container using the scripts/Makefile
start-db:
	@$(MAKE) -C scripts/docker

# Define targets for the Go server
.PHONY: start
start:
	@echo "Starting the Go server..."
	go run $(GO_SERVER)

.PHONY: build
build:
	@echo "Building the Go server..."
	go build -o server $(GO_SERVER)

.PHONY: clean
clean: clean-ui
	@echo "Cleaning up Go server build..."
	rm -f server

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  install-ui    - Install dependencies for the Next.js app"
	@echo "  start-ui      - Run Next.js in development mode"
	@echo "  build-ui      - Build the Next.js app"
	@echo "  start-ui-prod - Start the Next.js app in production mode"
	@echo "  lint-ui       - Run linter for the Next.js app"
	@echo "  clean-ui      - Clean up the Next.js app"
	@echo "  start         - Start the Go server"
	@echo "  build         - Build the Go server"
	@echo "  clean         - Clean up the Go server build"

# End of Makefile
