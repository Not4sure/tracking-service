
# Stop docker containers
.PHONY: stop
stop:
	@if podman-compose down>/dev/null; then \
		: ; \
	else \
		docker-compose down; \
	fi

# Start docker compose containers
.PHONY: start
start:
	@if podman-compose up -d>/dev/null; then \
		: ; \
	else \
		docker-compose up -d; \
	fi

.PHONY: start_db
start_db:
	@if podman-compose up -d db migrate adminer>/dev/null; then \
		: ; \
	else \
		docker-compose up -d db migrate adminer; \
	fi

.PHONY: build
build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Just run a application
.PHONY: run
run:
	@go run ./cmd/api/main.go

# Live Reload
.PHONY: watch
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Run unit tests
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v

.PHONY: codegen
sqlgen:
	@sqlc generate

.PHONY: migrateup
migrateup:
	@migrate -database "postgres://tracking_service:password@localhost:5432/tracking_service?sslmode=disable" -path ./internal/common/db/migrations up

.PHONY: migratedown
migratedown:
	@migrate -database "postgres://tracking_service:password@localhost:5432/tracking_service?sslmode=disable" -path ./internal/common/db/migrations down
