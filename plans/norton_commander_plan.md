# Projektplan: Commander-1 (Norton Commander Klon)

## Vision

Eine effiziente, tastaturgesteuerte Alternative zu grafischen Dateimanagern,
basierend auf dem klassischen Zwei-Panel-Layout des Norton Commanders.
Plattformübergreifend entwickelt für Linux und macOS mit Unterstützung für
verschiedene Architekturen (ARM, x86).

## Kernfunktionen (MVP)

- **Zwei-Panel-Ansicht:** Parallele Anzeige von zwei Verzeichnissen.
- **Navigation:** Schnelles Wechseln von Verzeichnissen via Tastatur.
- **Dateioperationen:** Kopieren (F5), Verschieben (F6), Löschen (F8),
  Umbenennen.
- **Vorschau/Edit:** Schnelle Ansicht (F3) und Editor-Integration (F4).
- **Plattformunterstützung:** Natives Verhalten auf Linux und macOS,
  Unterstützung für ARM64 und x86_64 Architekturen.

## Technologie-Stack

- **Sprache:** Go (Golang)
- **TUI Framework:** [Bubbletea](https://github.com/charmbracelet/bubbletea)
  (für das Elm-Architecture Pattern)
- **Komponenten:** [Bubbles](https://github.com/charmbracelet/bubbles)
  (für Listen, Inputs etc.)
- **Styling:** [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## Implementierungsplan

1. **Phase 1: Grundgerüst**
   - Go Projekt initialisieren.
   - Bubbletea Basis-Loop aufsetzen.
   - Layout für zwei Panels (links/rechts) erstellen.
2. **Phase 2: Dateisystem-Navigation**
   - Funktion zum Lesen von Verzeichnissen.
   - Navigation (Up/Down, Enter zum Öffnen, Backspace für zurück).
   - Fokus-Wechsel zwischen den Panels (Tab).
3. **Phase 3: Dateioperationen**
   - Implementierung von Kopieren, Verschieben und Löschen.
   - Fortschrittsanzeige für große Operationen.
4. **Phase 4: Plattform-Optimierung**
   - Linux-Unterstützung (ARM/x86).
   - macOS-Unterstützung.
   - plattformspezifische Pfade und Berechtigungen.
5. **Phase 5: Erweiterte Funktionen**
   - Handling von versteckten Dateien.
   - Integration von Standard-Editoren.
