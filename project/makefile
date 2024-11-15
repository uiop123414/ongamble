# Set directory paths
BACKEND_DIR = .././backend-service
LOGGER_DIR = .././logger-service
MAIL_DIR = .././mail-service
FRONTEND_DIR = .././ongambl

BACKEND_BINARY=backendApp
LOGGER_BINARY=loggerServiceApp
AUTH_BINARY=authApp

# Backend (Go) commands
.PHONY: backend-run
backend-run:
	@echo "Starting backend server..."
	cd $(BACKEND_DIR) && go run ./...

.PHONY: backend-build
backend-build:
	@echo "Building backend server..."
	cd $(BACKEND_DIR) && go build -o main ./...

.PHONY: backend-clean
backend-clean:
	@echo "Cleaning backend build files..."
	rm -f $(BACKEND_DIR)/main

# Frontend (ReactJS) commands
.PHONY: frontend-install
frontend-install:
	@echo "Installing frontend dependencies..."
	cd $(FRONTEND_DIR) && npm install

.PHONY: frontend-run
frontend-run:
	@echo "Starting frontend server..."
	cd $(FRONTEND_DIR) && npm start

.PHONY: frontend-build
frontend-build:
	@echo "Building frontend application..."
	cd $(FRONTEND_DIR) && npm run build

.PHONY: frontend-clean
frontend-clean:
	@echo "Cleaning frontend build files..."
	rm -rf $(FRONTEND_DIR)/build

# Commands to run both backend and frontend
.PHONY: run
run:
	@echo "Starting backend and frontend servers..."
	make backend-run & make frontend-run

.PHONY: build
build: backend-build frontend-build

.PHONY: clean
clean: backend-clean frontend-clean


## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_backend build_logger_service build_mail_service
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_backend: builds the backend binary as a windows executable
build_backend:
	@echo "Building backend application..."
	cd ${BACKEND_DIR} && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./backendApp ./cmd/api
	@echo "Done!"

## build_logger_service: builds the logger binary as a windows executable
build_logger_service:
	@echo "Building logger_service application..."
	cd ${LOGGER_DIR} && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./loggerApp ./cmd/api
	@echo "Done!"

## build_mail_service: builds the mail binary as a windows executable
build_mail_service:
	@echo "Building logger_service application..."
	cd ${MAIL_DIR} && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./mailApp ./cmd/api
	@echo "Done!"
