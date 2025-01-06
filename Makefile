# Default target
all: build test

# Build the Go application
build:
	@echo "Building..."
	@go build -o main ./...

# Run Docker containers
docker-run:
	@if docker compose up -d --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up -d --build; \
	fi

# Stop and remove Docker containers
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Run the Go application
run:
	@go run ./...

# Clean up the project
clean:
	@echo "Cleaning..."
	@rm -f main

# Run tests (you may need to define this if you have tests)
test:
	@echo "Running tests..."
	@go test ./...
