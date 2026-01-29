# Projekt-spezifische Rules & Guidelines (Commander-1)

Dieses Dokument definiert die Regeln für die Entwicklung des Commander-1 (Norton Commander Klon für macOS).

## 1. Technologie-Stack & Patterns
- **Sprache:** Go (Golang).
- **TUI Framework:** Bubbletea (Elm Architecture: Model, Update, View).
- **Styling:** Lip Gloss für alle UI-Komponenten.
- **Concurrency:** Nutze `tea.Cmd` für asynchrone Dateioperationen, um das UI nicht zu blockieren.

## 2. Code-Struktur
- `main.go`: Einstiegspunkt und Programm-Loop.
- `ui/`: Komponenten für das Interface (Panels, Listen, Footer).
- `fs/`: Logik für Dateisystem-Interaktionen (Lesen, Kopieren, Verschieben).
- `config/`: Benutzerkonfigurationen und Keybindings.

## 3. Test-Standards
- **Unit-Tests:** Logik in `fs/` muss durch `*_test.go` Dateien abgedeckt sein.
- **TUI-Tests:** Nutze die Test-Möglichkeiten von Bubbletea, um UI-Updates zu validieren.

## 4. macOS Spezifika
- Beachte macOS-spezifische Berechtigungen (Sandboxing/Full Disk Access).
- Versteckte Dateien (beginnend mit `.`) sollten standardmäßig ausblendbar sein.
- Nutze native macOS-Befehle via `os/exec`, falls Go-Standardbibliotheken nicht ausreichen (z.B. für Finder-spezifische Metadaten).

## 5. Kommunikation & Dokumentation
- Neue Features müssen zuerst in `plans/` skizziert werden.
- Wichtige Architektur-Entscheidungen werden als ADR in `docs/adr/` festgehalten.
