# Makefile

# Variables
DOCKER_COMPOSE_FILE = docker-compose.yaml
DOCKER_IMAGE_NAME = facial-recognition
GO_MAIN_FILE = ./src/cmd/server/main.go

# Test
.PHONY: test
test:
	echo "Running tests..."

# Run the server with Docker
.PHONY: all
all: build run

# Build the Docker image
.PHONY: build
build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Run the Docker container
.PHONY: run
run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Stop the Docker container
.PHONY: stop
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Run code locally
.PHONY: dev
dev:
	air