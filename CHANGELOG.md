# Changelog

Alle wichtigen Änderungen an diesem Projekt werden in dieser Datei dokumentiert.

Das Format basiert auf [Conventional Commits](https://www.conventionalcommits.org/) und verwendet [Semantic Versioning](https://semver.org/).

## [1.2.1] - 2026-01-31

### Fixed
- Korrigierte Tastenkürzel in der Man-Page entsprechend der tatsächlichen Implementierung

## [1.2.0] - 2026-01-31

### Added
- Man-page (manual) für min-commander Befehl
- Erweiterte Paketbeschreibung mit Hinweis auf Befehlsname

### Changed
- Konsistente Benennung aller Release-Artefakte zu 'min-commander-*'
- Verbesserte Dokumentation im Debian-Paket

### Fixed
- Fehlende Man-Page für installiertes Paket

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

- **Refactoring**: Interne Code-Struktur verbessert und modernisiert
- **CI-Fixes**: GitHub Actions Workflows stabilisiert und korrigiert

### Removed

- **Features**: Veraltete oder nicht mehr unterstützte Funktionen entfernt (Major Release aufgrund von Breaking Changes)

## [0.3.0] - 2026-01-31

### Features

- **Dateioperationen**: Kopieren (c), Verschieben (r), Löschen (d) für Dateien und Verzeichnisse
- **Verzeichnis-Rekursion**: CopyDir() und DeleteDir() für rekursives Kopieren/Löschen von Ordnern
- **Scrollbalken**: Vertikale Scrollbalken werden automatisch eingeblendet
- **Viewport-Steuerung**: PgUp/PgDn für schnelles Scrollen (10 Zeilen)
- **Automatisches Scrollen**: Viewport folgt dem Cursor automatisch
- **GitHub Actions CI**: Automatisierte Tests und Linting für alle Plattformen
- **GitHub Actions Release**: Automatische Release-Erstellung mit Binaries

### Bug Fixes

- Panels werden nach Dateioperationen korrekt neu geladen
- Temporäre Tastaturbelegung (c/r/d statt F5/F6/F8) um VSCode-Konflikte zu vermeiden

## [0.2.0] - 2026-01-29

### Features

- **Versteckte Dateien**: Toggle mit 'h' um versteckte Dateien anzuzeigen/auszublenden
- **Datei-Anzeige**: Datei-Viewer mit 'v' oder 'F3'
  - Textdateien: Zeilenweise Anzeige
  - Bilder: Öffnen mit externem Betrachter (xdg-open)
  - Binärdateien: Hexdump-Anzeige
- **Rekursive Dateisuche**: '/' für Wildcard-Suche (* und ?) mit Pfadangabe
  - Case-insensitive Matching
  - Eigene Suchergebnis-View
  - ENTER öffnet Dateien/Verzeichnisse aus Ergebnissen

### Changed

- Beide Panels werden bei Programmstart initialisiert
- Navigation überspringt versteckte Dateien wenn showHidden=false

## [0.1.0] - 2026-01-29

### Features

- Initiales Projekt-Setup mit Go und Bubbletea
- Zwei-Panel-Layout mit dynamischer Größenanpassung
- Dateisystem-Navigation (Pfeiltasten, Tab, Enter, Backspace)
- GitHub Repository Integration unter `sternrassler/commander-1`
- Projekt-spezifische Regeln und Richtlinien
- Linux-Unterstützung (x86_64, ARM64) und Cross-Compilation Makefile
- Architecture Decision Record (ADR-0001) für Plattform-Support
- Linting-Infrastruktur für Go und Markdown

### Changed

- README.md und Projektplan auf Cross-Plattform-Fokus aktualisiert
- Code-Refactoring zur Behebung von Linter-Warnungen (SA1019, ineffassign)
