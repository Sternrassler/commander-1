# Changelog

Alle wichtigen Änderungen an diesem Projekt werden in dieser Datei dokumentiert.

## [0.1.0] - 2026-01-29

### Hinzugefügt

- Initiales Projekt-Setup mit Go und Bubbletea.
- Zwei-Panel-Layout mit dynamischer Größenanpassung.
- Dateisystem-Navigation (Pfeiltasten, Tab, Enter, Backspace).
- GitHub Repository Integration unter `sternrassler/commander-1`.
- Projekt-spezifische Regeln und Richtlinien.
- Linux-Unterstützung (x86_64, ARM64) und Cross-Compilation Makefile.
- Architecture Decision Record (ADR-0001) für Plattform-Support.
- Linting-Infrastruktur für Go und Markdown.

### Geändert

- README.md und Projektplan auf Cross-Plattform-Fokus aktualisiert.
- Code-Refactoring zur Behebung von Linter-Warnungen (SA1019, ineffassign).

### Neu in Version 0.1.1

#### Features

- **Dateioperationen**: Kopieren (c), Verschieben (r), Löschen (d) für Dateien und Verzeichnisse.
- **Verzeichnis-Rekursion**: CopyDir() und DeleteDir() für rekursives Kopieren/Löschen von Ordnern.
- **Scrollbalken**: Vertikale Scrollbalken werden automatisch eingeblendet, wenn mehr Dateien als die Panel-Höhe vorhanden sind.
- **Viewport-Steuerung**: PgUp/PgDn für schnelles Scrollen (10 Zeilen).
- **Automatisches Scrollen**: Viewport folgt dem Cursor automatisch.

#### Fixes

- Panels werden nach Dateioperationen korrekt neu geladen.
- Temporäre Tastaturbelegung (c/r/d statt F5/F6/F8) um VSCode-Konflikte zu vermeiden.
