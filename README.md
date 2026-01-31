# Min Commander

Min Commander ist ein moderner, tastaturgesteuerter Terminal-Dateimanager als Alternative zum macOS Finder, inspiriert vom klassischen Norton Commander.

## Features

- **Zwei-Panel-Layout**: Effizientes Arbeiten in zwei Verzeichnissen gleichzeitig
- **TUI (Terminal User Interface)**: Schnell, leichtgewichtig und vollständig tastaturgesteuert
- **Cross-Platform**: Optimiert für macOS und Linux

### Dateioperationen

- **c**: Datei/Verzeichnis kopieren
- **r**: Datei/Verzeichnis verschieben
- **d**: Datei/Verzeichnis löschen

### Navigation

- **↑/↓**: Durch die Dateiliste navigieren
- **PgUp/PgDn**: Schnelles Scrollen (10 Zeilen)
- **Tab**: Zwischen linkem und rechtem Panel wechseln
- **Enter**: Verzeichnis öffnen
- **Backspace**: Zum übergeordneten Verzeichnis wechseln
- **h**: Versteckte Dateien anzeigen/ausblenden

### Dateibetrachter

- **v** oder **F3**: Datei anzeigen
  - Textdateien: Zeilenweise Anzeige
  - Bilder: Öffnen mit externem Betrachter
  - Binärdateien: Hexdump-Anzeige

### Dateisuche

- **/**: Wildcard-Suche (* und ?) mit Pfadangabe
  - Case-insensitive Matching
  - ENTER öffnet Dateien/Verzeichnisse aus Ergebnissen

## Installation

### Homebrew (macOS)

```bash
brew install sternrassler/tap/min-commander
```

### Direkter Download

Laden Sie die neueste Version von der [Release-Seite](https://github.com/sternrassler/commander-1/releases) herunter:

```bash
# macOS ARM64 (Apple Silicon)
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/min-commander-darwin-arm64 -o min-commander
chmod +x min-commander

# macOS x86_64 (Intel)
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/min-commander-darwin-amd64 -o min-commander
chmod +x min-commander

# Linux x86_64
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/min-commander-linux-amd64 -o min-commander
chmod +x min-commander

# Linux ARM64
curl -L https://github.com/sternrassler/commander-1/releases/latest/download/min-commander-linux-arm64 -o min-commander
chmod +x min-commander
```

### Von der Quelle bauen

```bash
git clone https://github.com/sternrassler/commander-1.git
cd commander-1
go build -o min-commander .
```

## Unterstützte Plattformen

- **Linux:** x86_64, ARM64 (aarch64)
- **macOS:** ARM64 (Apple Silicon), x86_64 (Intel)

## Build

Stellen Sie sicher, dass Go installiert ist.

### Build mit Make

Das Projekt enthält ein Makefile für Cross-Compilation:

```bash
# Alle Plattformen bauen
make all

# Spezifische Plattform
make linux-amd64
make linux-arm64
make darwin-amd64
make darwin-arm64

# Build-Artefakte bereinigen
make clean
```

Verfügbare Make Targets:

- `linux-amd64` (x86_64)
- `linux-arm64` (aarch64)
- `darwin-amd64` (macOS x86_64/Intel)
- `darwin-arm64` (macOS ARM64/Apple Silicon)
- `test` (Tests ausführen)
- `test-coverage` (Tests mit Coverage)
- `test-fs` (fs-tests mit Coverage)
- `test-integration` (Integration tests)
- `lint` (Code und Docs Linting)
- `lint-go` (Nur Go Code)
- `lint-docs` (Nur Docs)
- `install-lint` (Linting Tools installieren)

### Linting

Code-Qualität und Dokumentation werden mit Linting-Tools geprüft:

```bash
# Linting Tools installieren
make install-lint

# Alle Lints ausführen (Code + Docs)
make lint

# Nur Go Code
make lint-go

# Nur Docs
make lint-docs
```

## Steuerung

- **Pfeiltasten (↑/↓):** Durch die Dateiliste navigieren
- **Tab:** Zwischen linkem und rechtem Panel wechseln
- **Enter:** Verzeichnis öffnen
- **Backspace:** Zum übergeordneten Verzeichnis wechseln
- **q / Ctrl+C:** Beenden
- **c:** Kopieren
- **r:** Verschieben
- **d:** Löschen
- **h:** Versteckte Dateien umschalten
- **v / F3:** Datei anzeigen
- **/**: Dateisuche

## Tests und Coverage

```bash
# Alle Tests ausführen
make test

# Tests mit Coverage-Report
make test-coverage

# fs-tests (84.2% Coverage)
make test-fs

# Integration Tests
make test-integration
```

## Lizenz

Dieses Projekt ist unter der [MIT Lizenz](LICENSE) lizenziert.
