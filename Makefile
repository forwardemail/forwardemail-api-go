.PHONY: all build test lint fmt clean generate deps help

# Default target
all: deps lint test

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage report
test-coverage: test
	@echo "Generating coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Format code
fmt:
	@echo "Formatting code..."
	@gofumpt -l -w .

# Run code generation
generate:
	@echo "Running code generation..."
	@cd tools; go generate ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f coverage.out coverage.html
	@go clean -cache -testcache

# Install pre-commit hooks
hooks:
	@echo "Installing pre-commit hooks..."
	@pre-commit install

# Run pre-commit on all files
pre-commit:
	@echo "Running pre-commit..."
	@pre-commit run --all-files

# Verify module
verify:
	@echo "Verifying module..."
	@go mod verify

# Check for vulnerabilities
vuln:
	@echo "Checking for vulnerabilities..."
	@go run golang.org/x/vuln/cmd/govulncheck@latest ./...

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Run deps, lint, and test (default)"
	@echo "  deps         - Install dependencies"
	@echo "  test         - Run tests with race detection"
	@echo "  test-coverage - Run tests and generate HTML coverage report"
	@echo "  lint         - Run golangci-lint"
	@echo "  fmt          - Format code with gofumpt"
	@echo "  generate     - Run code generation"
	@echo "  clean        - Remove build artifacts"
	@echo "  hooks        - Install pre-commit hooks"
	@echo "  pre-commit   - Run pre-commit on all files"
	@echo "  verify       - Verify module dependencies"
	@echo "  vuln         - Check for vulnerabilities"
	@echo "  help         - Show this help message"

