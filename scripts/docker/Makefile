# scripts/Makefile

# Variables
DOCKER_IMAGE := postgres:latest
CONTAINER_NAME := videotube-db 
POSTGRES_USER := root
POSTGRES_PASSWORD := password 
POSTGRES_DB := test
POSTGRES_PORT := 5432

.PHONY: all check-docker pull-image run-container

# Default target
all: check-docker pull-image run-container

# Check if Docker is installed
check-docker:
	@command -v docker > /dev/null 2>&1 || { \
		echo >&2 "Docker is not installed. Please install Docker and try again."; \
		exit 1; \
	}

# Pull PostgreSQL Docker image if not present
pull-image:
	@if ! docker image inspect $(DOCKER_IMAGE) > /dev/null 2>&1; then \
		echo "Pulling PostgreSQL Docker image..."; \
		docker pull $(DOCKER_IMAGE); \
	fi

# Run PostgreSQL Docker container
run-container:
	@echo "Running PostgreSQL Docker container..."
	@docker run --name $(CONTAINER_NAME) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-d -p $(POSTGRES_PORT):5432 $(DOCKER_IMAGE)
	@echo "PostgreSQL container is running. You can connect to it on port $(POSTGRES_PORT)."
