# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Conventional Commits](https://www.conventionalcommits.org/)
and uses [Semantic Versioning](https://semver.org/).

## [2.1.1] - 2026-02-01

### Fixed

- **CopyDir**: Allow copying into existing directories
  - Removed restriction that prevented copying into existing destination
    directories
  - Now only checks for file conflicts, not directory existence
  - Enables copying directories into existing target locations (e.g., copying
    `.kilocode` into an existing project directory)

### Changed

- **Tests**: Updated `TestCopyDir_DestinationExists` to reflect new behavior

## [2.1.0] - 2026-02-01

### Added

- **Recursive File Operations**: Full support for recursive operations on
  directories
  - Copy directories recursively with all subdirectories and files
  - Move directories recursively across partitions
  - Delete directories recursively with all contents

### Changed

- **File Operations**: Refactored to use fs package functions consistently
  - `fs.CopyDir()` for recursive directory copying
  - `fs.DeleteDir()` for recursive directory deletion
  - `fs.Move()` for moving files and directories (with cross-partition support)
- **Code Quality**: Removed duplicate `copyFile()` function in favor of
  `fs.Copy()`
- **Tests**: Updated all tests to use fs package functions

### Fixed

- **Directory Deletion**: Fixed "directory not empty" error when deleting
  directories
- **Directory Operations**: All file operations now properly handle both files
  and directories

## [2.0.0] - 2026-01-31

### Changed

- **Documentation**: All documentation and code comments are now in English
  only
- **Project Rules**: Updated .kilocode/rules.md with TDD and Clean Code
  methodology
- **Workflow**: Added branch protection rule (no direct push to main)
- **Infrastructure**: Added .gitignore to prevent package-lock.json issues
- **Linting**: Fixed all markdown linting errors

### Removed

- Obsolete plan files (deb_package_plan.md, norton_commander_plan.md)
- package-lock.json (now in .gitignore)

### BREAKING CHANGES

- All documentation is now in English only. German documentation has been
  removed.

## [1.2.1] - 2026-01-31

### Fixed

- Corrected keyboard shortcuts in man page according to actual implementation

## [1.2.0] - 2026-01-31

### Added

- Man-page (manual) for min-commander command
- Extended package description with reference to command name

### Changed

- Consistent naming of all release artifacts to 'min-commander-*'
- Improved documentation in Debian package

### Removed

- Missing man page for installed package

## [1.1.0] - 2026-01-31

### Added

- Debian package (.deb) support for Linux amd64 and arm64 architectures
- nfpm configuration for package generation
- Automated .deb package building in release workflow

### Changed

- Enhanced release workflow to include .deb packages in GitHub releases
- Improved checksums generation to include all release artifacts

### Fixed

- Corrected tag_name reference in release workflow to use proper Git tag format

## [1.0.0] - 2026-01-31

### Changed

- **Refactoring**: Improved and modernized internal code structure
- **CI-Fixes**: Stabilized and corrected GitHub Actions workflows

### Removed

- **Features**: Removed deprecated or no longer supported features
  (Major Release due to Breaking Changes)

## [0.3.0] - 2026-01-31

### Features

- **File Operations**: Copy (c), Move (r), Delete (d) for files and
  directories
- **Directory Recursion**: CopyDir() and DeleteDir() for recursive
  copying/deleting of folders
- **Scrollbars**: Vertical scrollbars are automatically displayed
- **Viewport Control**: PgUp/PgDn for fast scrolling (10 lines)
- **Automatic Scrolling**: Viewport follows cursor automatically
- **GitHub Actions CI**: Automated tests and linting for all platforms
- **GitHub Actions Release**: Automatic release creation with binaries

### Bug Fixes

- Panels are correctly reloaded after file operations
- Temporary key bindings (c/r/d instead of F5/F6/F8) to avoid VSCode
  conflicts

## [0.2.0] - 2026-01-29

### Features

- **Hidden Files**: Toggle with 'h' to show/hide hidden files
- **File Viewer**: File viewer with 'v' or 'F3'
  - Text files: Line-by-line display
  - Images: Open with external viewer (xdg-open)
  - Binary files: Hexdump display
- **Recursive File Search**: '/' for wildcard search (* and ?) with path
  specification
  - Case-insensitive matching
  - Dedicated search results view
  - ENTER opens files/directories from results

### Changed

- Both panels are initialized at program start
- Navigation skips hidden files when showHidden=false

## [0.1.0] - 2026-01-29

### Features

- Initial project setup with Go and Bubbletea
- Dual-panel layout with dynamic resizing
- File system navigation (arrow keys, Tab, Enter, Backspace)
- GitHub repository integration under `sternrassler/commander-1`
- Project-specific rules and guidelines
- Linux support (x86_64, ARM64) and cross-compilation Makefile
- Architecture Decision Record (ADR-0001) for platform support
- Linting infrastructure for Go and Markdown

### Changed

- Updated README.md and project plan to cross-platform focus
- Code refactoring to fix linter warnings (SA1019, ineffassign)
