# Commander-1

Commander-1 ist eine moderne, tastaturgesteuerte Terminal-Alternative zum
macOS Finder, inspiriert vom klassischen Norton Commander.

## Features

- **Zwei-Panel-Layout:** Effizientes Arbeiten in zwei Verzeichnissen
  gleichzeitig.
- **TUI (Terminal User Interface):** Schnell, leichtgewichtig und komplett
  tastaturgesteuert.
- **Cross-Plattform:** Optimiert für macOS und Linux.

## Unterstützte Plattformen

- **Linux:** x86_64, ARM64 (aarch64)
- **macOS:** ARM64 (Apple Silicon), x86_64 (Intel)

## Build

Stelle sicher, dass Go installiert ist.

### Build mit Make

Das Projekt enthält ein Makefile für die Cross-Compilation:

```bash
# Alle Plattformen bauen
make all

# Spezifische Plattform
make linux-amd64
make linux-arm64
make darwin-amd64
make darwin-arm64

# Build-Artefakte aufräumen
make clean
```

Unterstützte Make-Targets:

- `linux-amd64` (x86_64)
- `linux-arm64` (aarch64)
- `darwin-amd64` (macOS x86_64/Intel)
- `darwin-arm64` (macOS ARM64/Apple Silicon)
- `lint` (Code und Docs linten)
- `lint-go` (nur Go-Code linten)
- `lint-docs` (nur Docs linten)
- `install-lint` (Linting-Tools installieren)

### Linting

Code-Qualität und Dokumentation werden mit Linting-Tools geprüft:

```bash
# Linting-Tools installieren
make install-lint

# Alle Lints ausführen (Code + Docs)
make lint

# Nur Go-Code linten
make lint-go

# Nur Dokumentation linten
make lint-docs
```

### Lokaler Build (alle Plattformen)

Mit Go kannst du direkt für deine aktuelle Plattform bauen:

```bash
go run main.go
```

Oder für macOS kompilieren:

```bash
# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o commander-1-darwin-arm64 .

# macOS x86_64 (Intel)
GOOS=darwin GOARCH=amd64 go build -o commander-1-darwin-amd64 .
```

## Steuerung

- **Pfeiltasten (↑/↓):** Navigieren durch die Dateiliste.
- **Tab / Pfeiltasten (←/→):** Wechseln zwischen linkem und rechtem Panel.
- **Enter:** Verzeichnis öffnen.
- **Backspace:** In das übergeordnete Verzeichnis wechseln.
- **q / Ctrl+C:** Beenden.

## Lizenz

Dieses Projekt steht unter der [MIT Lizenz](LICENSE).
