# Contributing to Min Commander

Thank you for your interest in Min Commander! We welcome contributions from
the community.

## How to Contribute

### 1. Reporting Issues

If you find a bug or have a feature request:

- Search first to see if the issue already exists
- Create a new issue with a clear description
- Include steps to reproduce (for bugs)
- Describe the expected behavior

### 2. Submitting Pull Requests

#### Preparation

1. Fork the repository
2. Clone your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/commander-1.git
   cd commander-1
   ```

3. Create a new branch for your changes:

   ```bash
   git checkout -b feature/your-feature-name
   ```

#### Code Standards

- **Go Version**: The project uses Go 1.21 or higher
- **Code Formatting**: Use `go fmt`
- **Linting**: Run `make lint-go` before committing
- **Tests**: All tests must pass (`make test`)

#### Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/)
specification:

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:**

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style**: Formatting changes (no logic changes)
- `refactor`: Code restructuring without behavior change
- `test`: Adding or modifying tests
- `chore**: Maintenance tasks

**Examples:**

```text
feat(fs): Add recursive directory deletion support
fix: Correct cursor movement when hiding files
docs: Update installation instructions
```

#### Testing

- Add tests for new features
- Ensure all tests pass:

  ```bash
  make test
  ```

- Coverage target for fs package: at least 80%

  ```bash
  make test-fs
  ```

#### Creating a Pull Request

1. Run all tests and lints
2. Push your branch
3. Create a pull request via GitHub
4. Describe your changes in detail
5. Link relevant issues

### 3. Review Process

- All pull requests are reviewed by maintainers
- Changes may be requested
- Please respond to feedback promptly

## Workflow

### Issue-Driven Development

1. **Create an Issue**: Before making changes, create an issue that describes
   your change request.
   - Choose a meaningful title
   - Describe the problem or desired feature
   - Explain the expected benefit

2. **Use Labels**: Add appropriate labels to the issue:
   - `feature`: For new feature requests
   - `bug`: For bug reports and fixes

3. **Work in your own Branch**: Create a branch based on the issue:

   ```bash
   git checkout -b feature/issue-123-description   # for Features
   git checkout -b fix/issue-456-description       # for Bugfixes
   ```

4. **Create a Pull Request**: After completion, create a PR against the
   `main` branch.
   - PR title should reference the issue
     (e.g. "feat: Add recursive deletion (#123)")
   - Link the corresponding issue in the PR

### Clean Code Methodology

Follow the Clean Code principles by Robert C. Martin:

- **SOLID Principles**: Single Responsibility, Open/Closed, Liskov
  Substitution, Interface Segregation, Dependency Inversion
- **KISS**: "Keep It Simple, Stupid" - Keep the code as simple as possible
- **DRY**: "Don't Repeat Yourself" - Avoid code duplication
- **YAGNI**: "You Aren't Gonna Need It" - Implement only what is really
  needed
- **Meaningful Names**: Use descriptive names for variables, functions, and types
- **Small Functions**: Functions should perform one task and be at most
  20-30 lines long
- **Comments**: Comment only "why", not "what" - the code should be
  self-explanatory

## Project Structure

```text
commander-1/
├── main.go           # Main application and Model
├── fs/
│   ├── fs.go         # File system functions
│   └── fs_test.go    # Tests for fs functions
├── main_test.go      # Unit tests for main functions
├── integration_test.go # Integration tests
├── Makefile          # Build and test targets
├── README.md         # Project documentation
└── CONTRIBUTING.md   # This file
```

## Coding Guidelines

### Go

- Use descriptive variable names
- Comment public functions
- Handle errors explicitly (don't use `_` for errors)
- Use `t.TempDir()` for temporary test directories

### Tests

- Unit tests for all public functions
- Integration tests for complex workflows
- Avoid flaky tests (non-deterministic results)

## Questions?

For questions you can:

- Create an issue
- Use Discussions
- Contact maintainers directly

Thank you for your support!
