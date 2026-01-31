# Project-Specific Rules & Guidelines (Commander-1)

This document defines the rules for developing Commander-1 (Norton Commander clone for macOS).

## 1. Technology Stack & Patterns
- **Language:** Go (Golang) 1.21 or higher
- **TUI Framework:** Bubbletea (Elm Architecture: Model, Update, View)
- **Styling:** Lip Gloss for all UI components
- **Concurrency:** Use `tea.Cmd` for asynchronous file operations to avoid blocking the UI

## 2. Code Structure
- `main.go`: Entry point and program loop
- `ui/`: Components for the interface (panels, lists, footer)
- `fs/`: Logic for file system interactions (reading, copying, moving)
- `config/`: User configurations and keybindings

## 3. Development Methodology

### Test-Driven Development (TDD)
- Write tests before implementing features
- Follow the Red-Green-Refactor cycle:
  1. Write a failing test (Red)
  2. Implement minimal code to pass the test (Green)
  3. Refactor while keeping tests green
- Aim for >80% code coverage (especially in `fs/` package)

### Clean Code Principles
Follow Robert C. Martin's Clean Code methodology:

- **SOLID Principles:** Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion
- **KISS:** "Keep It Simple, Stupid" - Keep the code as simple as possible
- **DRY:** "Don't Repeat Yourself" - Avoid code duplication
- **YAGNI:** "You Aren't Gonna Need It" - Implement only what is really needed
- **Meaningful Names:** Use descriptive names for variables, functions, and types
- **Small Functions:** Functions should perform one task and be at most 20-30 lines long
- **Self-Documenting Code:** Code should be readable without comments
- **Comments:** Comment only "why", not "what" - explain intent, not implementation

## 4. Test Standards
- **Unit Tests:** Logic in `fs/` must be covered by `*_test.go` files
- **TUI Tests:** Use Bubbletea's testing capabilities to validate UI updates
- **Integration Tests:** Test complex workflows in `integration_test.go`
- **Test Coverage:** Minimum 80% for `fs/` package
- **Test Naming:** Use descriptive test names that explain what is being tested
- **Test Isolation:** Use `t.TempDir()` for temporary test directories
- **No Flaky Tests:** Avoid non-deterministic test results

## 5. Platform Support
As defined in [ADR-0001](docs/adr/0001-linux-support.md):

### Supported Platforms
- **Linux:** x86_64, ARM64 (aarch64)
- **macOS:** ARM64 (Apple Silicon), x86_64 (Intel)

### Unsupported Platforms
- **Windows:** Not supported and not planned for future support

### macOS Specifics
- Consider macOS-specific permissions (Sandboxing/Full Disk Access)
- Hidden files (starting with `.`) should be hideable by default
- Use native macOS commands via `os/exec` if Go standard libraries are insufficient (e.g., for Finder-specific metadata)

## 6. Communication & Documentation
- **Communication:** Communicate with the developer in German
- **Documentation:** Document EVERYTHING in English (code comments, README, ADRs, CONTRIBUTING.md, internal rules)
- **Architecture Decisions:** Important architecture decisions are recorded as ADRs in `docs/adr/`
- **Changelog:** Follow [Conventional Commits](https://www.conventionalcommits.org/) and [Semantic Versioning](https://semver.org/)

## 7. Contribution Workflow
As defined in [CONTRIBUTING.md](CONTRIBUTING.md):

### Issue-Driven Development
1. Create an issue before making changes
2. Use appropriate labels (`feature`, `bug`)
3. Work in feature branches: `feature/issue-123-description` or `fix/issue-456-description`
4. Create PR against `main` branch with issue reference
5. **NEVER push directly to `main`** - all changes must go through PR review process

### Code Quality
- **Formatting:** Use `go fmt` before committing
- **Linting:** Run `make lint-go` before committing
- **Tests:** All tests must pass (`make test`)
- **Error Handling:** Handle errors explicitly (don't use `_` for errors)
- **Public Functions:** Comment all public functions

### Commit Messages
Follow Conventional Commits:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

**Example:**
```
feat(fs): Add recursive directory deletion support
fix: Correct cursor movement when hiding files
docs: Update installation instructions
```

## 8. Code Review
- All pull requests are reviewed by maintainers
- Changes may be requested
- Respond to feedback promptly
- Ensure all CI checks pass before requesting review
