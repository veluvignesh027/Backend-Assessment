docker-run:
    @if docker compose up -d --build 2>/dev/null; then \
        : ; \
    else \
        echo "Falling back to Docker Compose V1"; \
        docker-compose up -d --build; \
    fi

docker-down:
	@if docker compose down 2>/dev/null; then \
        : ; \
    else \
        echo "Falling back to Docker Compose V1"; \
        docker-compose down; \
    fi

run:
    @go run ./...

clean:
    @echo "Cleaning..."
    @rm -f main