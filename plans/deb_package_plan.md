# Plan zur Generierung von .deb-Paketen für Min Commander

Dieses Dokument beschreibt den Plan zur Einführung von Debian-Paketen (.deb) für das Projekt `commander-1` (Min Commander).

## 1. Gewählter Ansatz: `nfpm`

Nach Abwägung der Optionen (`GoReleaser`, `nfpm`, `dpkg-deb`) wird **nfpm** empfohlen.

### Begründung
*   **GoReleaser:** Wäre eine Komplettlösung, erfordert aber eine größere Umstellung des bestehenden, funktionierenden GitHub Actions Workflows.
*   **dpkg-deb:** Erfordert eine komplexe Verzeichnisstruktur und manuelle Control-Dateien, was fehleranfällig ist.
*   **nfpm:** Bietet die perfekte Balance. Es ist ein leichtgewichtiges Tool, das eine einfache YAML-Konfiguration nutzt und sich nahtlos in den bestehenden Workflow integrieren lässt, ohne diesen komplett zu ersetzen.

## 2. Paket-Metadaten

Basierend auf den Projektdateien werden folgende Metadaten verwendet:

*   **Name:** `min-commander` (entspricht dem Binärnamen im Makefile)
*   **Version:** Dynamisch aus der Datei `VERSION` oder dem Git-Tag.
*   **Beschreibung:** "Ein moderner, tastaturgesteuerter Terminal-Dateimanager als Alternative zum macOS Finder, inspiriert vom klassischen Norton Commander." (aus `README.md`)
*   **Maintainer:** (Platzhalter, z.B. `Sternrassler <support@example.com>`)
*   **Homepage:** `https://github.com/sternrassler/commander-1`
*   **Lizenz:** `MIT`
*   **Sektion:** `utils`
*   **Priorität:** `optional`

## 3. Technische Details

### Zielarchitekturen
*   `amd64` (x86_64)
*   `arm64` (aarch64)

### Installationspfad
Die Binärdatei wird unter `/usr/bin/min-commander` installiert.

### Konfigurationsdatei (`nfpm.yaml`)
Es wird eine `nfpm.yaml` im Wurzelverzeichnis erstellt:

```yaml
name: "min-commander"
arch: "amd64"
platform: "linux"
version: "${VERSION}"
section: "utils"
priority: "optional"
maintainer: "Sternrassler"
description: "Modern keyboard-driven terminal file manager"
homepage: "https://github.com/sternrassler/commander-1"
license: "MIT"
contents:
  - src: "min-commander-linux-amd64"
    dst: "/usr/bin/min-commander"
```

## 4. Änderungen am Workflow (`.github/workflows/release.yml`)

Der Release-Workflow wird um folgende Schritte erweitert:

1.  **Installation von nfpm:** Verwendung der offiziellen Action oder direkter Download.
2.  **Paketbau:**
    *   Nach dem `go build` für Linux AMD64/ARM64 wird `nfpm pkg --target min-commander_${VERSION}_amd64.deb` ausgeführt.
3.  **Artefakt-Upload:** Die generierten `.deb` Dateien werden zu den GitHub Release Assets hinzugefügt.

## 5. Implementierungsschritte (Todo)

1.  [ ] `nfpm.yaml` Konfigurationsdatei erstellen.
2.  [ ] Makefile erweitern, um lokale Paket-Builds zu unterstützen (optional, aber empfohlen).
3.  [ ] `.github/workflows/release.yml` anpassen:
    *   `nfpm` Setup-Schritt hinzufügen.
    *   Paket-Generierung nach den Build-Schritten einfügen.
    *   `.deb` Dateien zur Liste der hochzuladenden Dateien in `softprops/action-gh-release` hinzufügen.
    *   Checksummen-Generierung um `.deb` Dateien erweitern.
