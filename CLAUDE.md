# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Min Commander — a Norton Commander–inspired, keyboard-driven dual-panel terminal file manager for Linux and macOS. Written in Go using the Bubbletea TUI framework (Elm architecture) and Lip Gloss for styling. Windows is explicitly out of scope (see `docs/adr/0001-linux-support.md`).

TUI framework convention: long-running file-system work must go through `tea.Cmd` so the UI stays responsive — never call blocking `fs.*` operations directly from `Update`.

Module path: `github.com/karstenflache/commander-1`. Go 1.24 (toolchain pinned in `go.mod`); CONTRIBUTING.md still references 1.21+ as the minimum.

## Architecture

Two-package layout:

- `main.go` — entire TUI lives here: `model` (Bubbletea Model), `panel` struct (two of them in `model.panels`), `Init`/`Update`/`View`, key handling, async file ops dispatched as `tea.Cmd`s returning typed messages (e.g. `readDirMsg`). All rendering uses Lip Gloss styles defined at the top of the file.
- `fs/fs.go` — pure file-system operations (`ReadDir`, `Copy`, `Move`, `Delete`, recursive variants). No TUI dependencies; this is the package with the 80%+ coverage target.
- `main_test.go` — unit tests for the model/update logic.
- `integration_test.go` — end-to-end workflows (run with `make test-integration`, filter `TestIntegration`).

## Common commands

```bash
make test                  # all tests
make test-fs               # fs package only, with coverage (target ≥80%)
make test-integration      # only TestIntegration* in repo root
make lint                  # lint-go + lint-docs
make install-lint          # bootstrap golangci-lint + markdownlint
make all                   # cross-compile all four target binaries
make package-linux-amd64   # build .deb via nfpm (reads VERSION file)
```

Run a single Go test: `go test -v -run TestName ./fs` (or `.` for root package).

Releases: bump the `VERSION` file (plain text, e.g. `2.0.1`); `nfpm.yaml` reads it via `$VERSION`. Update `CHANGELOG.md` following Keep-a-Changelog style.

## Conventions worth knowing

- **Commit messages:** Conventional Commits (`feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`), optional scope like `feat(fs): ...`.
- **Branching:** Never push directly to `main`; work on `feature/issue-NNN-…` or `fix/issue-NNN-…` and open a PR referencing the issue.
- **Comments and docs in English**, even though the maintainer communicates in German. ADRs live in `docs/adr/`.
- **Error handling:** explicit — don't discard errors with `_`. Public functions get a doc comment.
- **Tests:** use `t.TempDir()` for any filesystem fixtures; keep tests deterministic.
