.PHONY: all clean linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 lint lint-go lint-docs install-lint

GOPATH_BIN := $(shell go env GOPATH)/bin

# Standard Build (alle Plattformen)
all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

# Linux AMD64 (x86_64)
linux-amd64:
	@echo "Building for linux-amd64..."
	GOOS=linux GOARCH=amd64 go build -o commander-1-linux-amd64 . && \
	@echo "✓ linux-amd64 build complete" || (echo "✗ linux-amd64 build failed" && exit 1)

# Linux ARM64 (aarch64)
linux-arm64:
	@echo "Building for linux-arm64..."
	GOOS=linux GOARCH=arm64 go build -o commander-1-linux-arm64 . && \
	@echo "✓ linux-arm64 build complete" || (echo "✗ linux-arm64 build failed" && exit 1)

# macOS AMD64 (Intel)
darwin-amd64:
	@echo "Building for darwin-amd64..."
	GOOS=darwin GOARCH=amd64 go build -o commander-1-darwin-amd64 . && \
	@echo "✓ darwin-amd64 build complete" || (echo "✗ darwin-amd64 build failed" && exit 1)

# macOS ARM64 (Apple Silicon)
darwin-arm64:
	@echo "Building for darwin-arm64..."
	GOOS=darwin GOARCH=arm64 go build -o commander-1-darwin-arm64 . && \
	@echo "✓ darwin-arm64 build complete" || (echo "✗ darwin-arm64 build failed" && exit 1)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f commander-1-linux-amd64 commander-1-linux-arm64 commander-1-darwin-amd64 commander-1-darwin-arm64
	@echo "✓ Clean complete"

# Linting

# Linting-Tools installieren
install-lint:
	@echo "Installing linting tools..."
	@command -v golangci-lint >/dev/null 2>&1 || (curl -sSfL https://raw.githubusercontent.com/golangci-lint/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_BIN) v1.64.5) || (echo "Warning: golangci-lint installation failed. Check if 'go' is in your PATH.")
	@command -v markdownlint >/dev/null 2>&1 || (npm install -g markdownlint-cli 2>/dev/null || npm install markdownlint-cli 2>/dev/null || echo "Warning: markdownlint installation failed")
	@echo "✓ Linting tools ready"

# Go-Code linten
lint-go:
	@echo "Linting Go code..."
	@PATH=$(PATH):$(GOPATH_BIN) golangci-lint run ./... || (echo "✗ golangci-lint failed or not found. Run 'make install-lint' first." && exit 1)
	@echo "✓ Go lint complete"

# Dokumentation linten
lint-docs:
	@echo "Linting documentation..."
	@PATH=$(PATH):./node_modules/.bin command -v markdownlint >/dev/null 2>&1 && markdownlint --ignore node_modules '**/*.md' || (echo "✗ markdownlint not found. Run 'make install-lint' first." && exit 1)
	@echo "✓ Docs lint complete"

# Alle Lints ausführen
lint: lint-go lint-docs
	@echo "✓ All linting complete"
