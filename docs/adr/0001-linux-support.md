# ADR-0001: Linux-Plattformunterstützung

## Status

Akzeptiert

## Kontext

Das Projekt soll als plattformübergreifende Anwendung entwickelt werden.
Eine wichtige Entscheidung betrifft die Unterstützung verschiedener
Betriebssysteme und Architekturen.

## Entscheidung

Das Projekt wird zukünftig die folgenden Plattformen unterstützen:

- **Linux** (x86_64, ARM64)
- **macOS** (ARM64, Apple Silicon)

### Nicht unterstützte Plattformen

- **Windows**: Das Projekt unterstützt Windows derzeit nicht und ist nicht
  für die zukünftige Unterstützung geplant.

## Begründung

1. **Linux-Fokus**: Als Open-Source-Projekt ist Linux die primäre
   Zielplattform für Entwicklung und Deployment.
2. **ARM-Unterstützung**: Mit dem Wachstum von ARM-Servern und
   Embedded-Geräten ist breite Architekturunterstützung wichtig.
3. **x86-Unterstützung**: Die Mehrheit der Desktop- und Serverumgebungen
   läuft noch auf x86, daher ist diese Architektur essenziell.
4. **Keine Windows-Unterstützung**: Die Entwicklungsressourcen werden auf
   Linux- und macOS-Kompatibilität konzentriert, um Stabilität und Qualität
   zu gewährleisten.

## Konsequenzen

### Positiv

- Klare Fokussierung auf Unix-ähnliche Umgebungen (Linux und macOS)
- Einfacheres Testen und Deployment
- Bessere Ressourcennutzung für Entwicklung

### Negativ

- Keine Möglichkeit für Windows-Benutzer, das Projekt direkt zu nutzen
- Potentieller Ausschluss eines Teils der potenziellen Nutzerbasis

## Referenzen

- Projektphilosophie: Open Source und Linux-freundlich
- Ressourcenbeschränkungen erfordern Priorisierung
