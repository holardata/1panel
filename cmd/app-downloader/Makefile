BINARY_NAME=app-downloader
BUILD_DIR=bin

.PHONY: all clean build-mac build-linux-amd64 build-linux-arm64

all: clean build-mac build-linux-amd64 build-linux-arm64

build-mac:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go

build-linux-amd64:
	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go

build-linux-arm64:
	@echo "Building for Linux (arm64)..."
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 main.go

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
