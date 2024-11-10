# Set directory paths
BACKEND_DIR = ./backend
FRONTEND_DIR = ./ongambl
ONGAMBL_DB_DSN=-postgres://postgres:Qzpm231414@localhost/ongambl?sslmode=disable

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

.PHONY: run
run:
	@echo "Starting backend and frontend servers..."
	powershell -Command "Start-Process cmd -ArgumentList '/c make backend-runmm'"
	powershell -Command "Start-Process cmd -ArgumentList '/c make frontend-run'"

.PHONY: build
build: backend-build frontend-build

.PHONY: clean
clean: backend-clean frontend-clean

.PHONY: stop
stop:
	@echo "Stopping backend and frontend servers..."
	taskkill /F /IM go.exe /T
	taskkill /F /IM node.exe /T