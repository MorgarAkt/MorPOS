BINARY_NAME := MorPOS
BIN_DIR := bin
SOURCE_DIRS := ./...
DB_PATH := database.db
MIGRATIONS_DIR := ./sql/migrations
GOOSE := ~/go/bin/goose

.PHONY: build test clean run setup migrate-up migrate-down

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/main.go

test:
	go test $(SOURCE_DIRS)

clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

run: build
	$(BIN_DIR)/$(BINARY_NAME)

setup:
	$(GOOSE) -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up

migrate-up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) down
