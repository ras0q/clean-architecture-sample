SHELL      := /bin/bash
LOCAL_PORT := 8081

# Run on Docker
all: stop up logs
stop:
	@docker compose stop
up:
	@docker compose up --build -d
logs:
	@docker compose logs -f

# Run locally
.PHONY: dev
local:
	@docker compose up db -d --no-recreate
	@PORT=$(LOCAL_PORT) go run github.com/cosmtrek/air@v1.27.3

# Test go-files
.PHONY: test
test:
	@docker compose exec app go test -cover -shuffle=on -v ./...

# Generate templates
gen:
	@go generate ./...
