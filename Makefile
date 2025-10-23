
.PHONY: build
build:
	@echo "Building..."
	@go build -o main cmd/app/main.go

# Start docker compose containers
.PHONY: start
start:
	@if podman-compose up -d>/dev/null; then \
		: ; \
	else \
		docker-compose up -d; \
	fi

# Just run a application
.PHONY: run
run:
	@go run ./cmd/app/main.go

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

