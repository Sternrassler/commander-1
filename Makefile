.PHONY: all clean linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 package-linux-amd64 install-darwin-arm64 lint lint-go lint-docs install-lint test test-coverage test-fs test-integration

GOPATH_BIN := $(shell go env GOPATH)/bin

# Standard Build (alle Plattformen)
all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

# Linux AMD64 (x86_64)
linux-amd64:
	@echo "Building for linux-amd64..."
	GOOS=linux GOARCH=amd64 go build -o min-commander-linux-amd64 . && \
	@echo "✓ linux-amd64 build complete" || (echo "✗ linux-amd64 build failed" && exit 1)

# Linux ARM64 (aarch64)
linux-arm64:
	@echo "Building for linux-arm64..."
	GOOS=linux GOARCH=arm64 go build -o min-commander-linux-arm64 . && \
	@echo "✓ linux-arm64 build complete" || (echo "✗ linux-arm64 build failed" && exit 1)

# macOS AMD64 (Intel)
darwin-amd64:
	@echo "Building for darwin-amd64..."
	GOOS=darwin GOARCH=amd64 go build -o min-commander-darwin-amd64 . && \
	@echo "✓ darwin-amd64 build complete" || (echo "✗ darwin-amd64 build failed" && exit 1)

# macOS ARM64 (Apple Silicon)
darwin-arm64:
	@echo "Building for darwin-arm64..."
	GOOS=darwin GOARCH=arm64 go build -o min-commander-darwin-arm64 . && \
	@echo "✓ darwin-arm64 build complete" || (echo "✗ darwin-arm64 build failed" && exit 1)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f min-commander-linux-amd64 min-commander-linux-arm64 min-commander-darwin-amd64 min-commander-darwin-arm64 coverage.out
	@echo "✓ Clean complete"

# Tests ausführen
test:
	@echo "Running all tests..."
	go test -v ./...
	@echo "✓ All tests complete"

# Tests mit Coverage ausführen
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	@echo "✓ Coverage report complete"

# Nur fs-Tests ausführen (Coverage-Ziel: 80%+)
test-fs:
	@echo "Running fs package tests..."
	go test -v -coverprofile=coverage-fs.out ./fs
	go tool cover -func=coverage-fs.out
	@echo "✓ fs tests complete"

# Integrationstests ausführen
test-integration:
	@echo "Running integration tests..."
	go test -v -run TestIntegration .
	@echo "✓ Integration tests complete"

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

# Pakete erstellen

# Linux AMD64 DEB-Paket
package-linux-amd64: linux-amd64
	@echo "Creating Linux AMD64 DEB package..."
	@command -v nfpm >/dev/null 2>&1 || (echo "✗ nfpm not found. Install with: go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest" && exit 1)
	VERSION=$$(cat VERSION) nfpm pkg --packager deb --config nfpm.yaml --target min-commander-linux-amd64.deb
	@echo "✓ Linux AMD64 DEB package created: min-commander-linux-amd64.deb"

# macOS ARM64 Installation (für M1/M2/M3 Macs)
install-darwin-arm64: darwin-arm64
	@echo "Installing macOS ARM64 binary..."
	sudo cp min-commander-darwin-arm64 /usr/local/bin/min-commander
	sudo chmod +x /usr/local/bin/min-commander
	@echo "✓ macOS ARM64 binary installed to /usr/local/bin/min-commander"
	@echo ""
	@echo "Programm starten mit: min-commander"
