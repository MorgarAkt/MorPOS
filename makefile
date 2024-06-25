BINARY_NAME := MorPOS
BIN_DIR := bin
SOURCE_DIRS := ./...

.PHONY: build test clean run

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/main.go

test:
	go test $(SOURCE_DIRS)

clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

run: build
	$(BIN_DIR)/$(BINARY_NAME)
