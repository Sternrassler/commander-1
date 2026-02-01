# Min Commander

Min Commander is a modern, keyboard-driven terminal file manager as an
alternative to macOS Finder, inspired by the classic Norton Commander.

## Features

- **Dual-Panel Layout**: Efficient work in two directories simultaneously
- **TUI (Terminal User Interface)**: Fast, lightweight and fully keyboard-driven
- **Cross-Platform**: Optimized for macOS and Linux

### File Operations

- **c**: Copy file/directory (recursive for directories)
- **r**: Move file/directory (recursive for directories, works across partitions)
- **d**: Delete file/directory (recursive for directories)

All file operations automatically detect whether the selected item is a file or
directory and handle it appropriately. Directories are processed recursively
with all their contents.

### Navigation

- **↑/↓**: Navigate through file list
- **PgUp/PgDn**: Fast scrolling (10 lines)
- **Tab**: Switch between left and right panel
- **Enter**: Open directory
- **Backspace**: Go to parent directory
- **h**: Show/hide hidden files

### File Viewer

- **v** or **F3**: View file
  - Text files: Line-by-line display
  - Images: Open with external viewer
  - Binary files: Hexdump display

### File Search

- **/**: Wildcard search (* and ?) with path specification
  - Case-insensitive matching
  - ENTER opens files/directories from results

## Installation

### Homebrew (macOS)

```bash
brew install sternrassler/tap/min-commander
```

### Direct Download

Download the latest version from the [Release page](https://github.com/sternrassler/commander-1/releases):

#### macOS

```bash
# macOS ARM64 (Apple Silicon M1/M2/M3)
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander-darwin-arm64.gz -o min-commander.gz
gunzip min-commander.gz
chmod +x min-commander
sudo mv min-commander /usr/local/bin/

# macOS x86_64 (Intel)
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander-darwin-amd64.gz -o min-commander.gz
gunzip min-commander.gz
chmod +x min-commander
sudo mv min-commander /usr/local/bin/
```

#### Linux

```bash
# Linux x86_64 - Binary
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander-linux-amd64.gz -o min-commander.gz
gunzip min-commander.gz
chmod +x min-commander
sudo mv min-commander /usr/local/bin/

# Linux x86_64 - DEB Package
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander_VERSION_amd64.deb -o min-commander.deb
sudo dpkg -i min-commander.deb

# Linux ARM64 - Binary
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander-linux-arm64.gz -o min-commander.gz
gunzip min-commander.gz
chmod +x min-commander
sudo mv min-commander /usr/local/bin/

# Linux ARM64 - DEB Package
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/\
min-commander_VERSION_arm64.deb -o min-commander.deb
sudo dpkg -i min-commander.deb
```

**Hinweis für DEB-Pakete:** Ersetze `VERSION` mit der aktuellen Versionsnummer (z.B. `2.0.1`).

### Build from Source

```bash
git clone https://github.com/sternrassler/commander-1.git
cd commander-1
go build -o min-commander .
```

## Supported Platforms

- **Linux:** x86_64, ARM64 (aarch64)
- **macOS:** ARM64 (Apple Silicon), x86_64 (Intel)

## Build

Make sure Go is installed.

### Build with Make

The project includes a Makefile for cross-compilation:

```bash
# Build all platforms
make all

# Specific platform
make linux-amd64
make linux-arm64
make darwin-amd64
make darwin-arm64

# Clean build artifacts
make clean
```

Available Make targets:

- `linux-amd64` (x86_64)
- `linux-arm64` (aarch64)
- `darwin-amd64` (macOS x86_64/Intel)
- `darwin-arm64` (macOS ARM64/Apple Silicon)
- `package-linux-amd64` (Create DEB package for Linux AMD64)
- `package-darwin-arm64` (Create PKG package for macOS ARM64)
- `test` (Run tests)
- `test-coverage` (Tests with coverage)
- `test-fs` (fs-tests with coverage)
- `test-integration` (Integration tests)
- `lint` (Code and docs linting)
- `lint-go` (Go code only)
- `lint-docs` (Docs only)
- `install-lint` (Install linting tools)

### Linting

Code quality and documentation are checked with linting tools:

```bash
# Install linting tools
make install-lint

# Run all lints (code + docs)
make lint

# Go code only
make lint-go

# Docs only
make lint-docs
```

## Controls

- **Arrow keys (↑/↓):** Navigate through file list
- **Tab:** Switch between left and right panel
- **Enter:** Open directory
- **Backspace:** Go to parent directory
- **q / Ctrl+C:** Quit
- **c:** Copy
- **r:** Move
- **d:** Delete
- **h:** Toggle hidden files
- **v / F3:** View file
- **/**: File search

## Tests and Coverage

```bash
# Run all tests
make test

# Tests with coverage report
make test-coverage

# fs-tests (84.2% coverage)
make test-fs

# Integration tests
make test-integration
```

## License

This project is licensed under the [MIT License](LICENSE).
