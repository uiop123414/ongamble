# Set directory paths
BACKEND_DIR = ./backend
FRONTEND_DIR = ./ongambl

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
