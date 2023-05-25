# Variables
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BINARY_NAME = fizzBuzzREST
BUILD_DIR = bin

.PHONY: all build test clean

all: build test

build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v .

test:
	cd ./ && $(GOCMD) mod tidy
	cd ./ && $(GOCMD) list -f '{{.Dir}}' ./... | xargs -L1 goimports -w
	cd ./ && $(GOTEST) -v ./... -count=1

clean:
	rm -rf $(BUILD_DIR)

run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v .
	$(BUILD_DIR)/$(BINARY_NAME) > $(BUILD_DIR)/output.log 2>&1
	@echo "Server logs:"
	@cat $(BUILD_DIR)/output.log
