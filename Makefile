# Makefile for the Web Crawler project

# Project name
PROJECT_NAME := webcrawler

# Directories
SRC_DIR := ./internal
BIN_DIR := ./bin

# Go commands
GO := go
GOFMT := $(GO) fmt
GOBUILD := $(GO) build
GOTEST := $(GO) test
GOCLEAN := $(GO) clean
GOVET := $(GO) vet
GOGET := $(GO) get

# Default target
all: build

# Format source files
fmt:
	$(GOFMT) ./...



# Build the project
build: fmt lint
	@echo "Building $(PROJECT_NAME)..."
	$(GOBUILD) -o $(BIN_DIR)/$(PROJECT_NAME) main.go

# Run tests
test: fmt lint
	$(GOTEST) -v $(SRC_DIR)/...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BIN_DIR)/*

# Install dependencies
deps:
	$(GOGET) -u ./...

# Run the application
run: build
	@echo "Running $(PROJECT_NAME)..."
	$(BIN_DIR)/$(PROJECT_NAME)

.PHONY: all fmt lint build test clean deps run
