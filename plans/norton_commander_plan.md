# Project Plan: Min Commander (Norton Commander Clone)

## Vision

An efficient, keyboard-driven alternative to graphical file managers,
based on the classic two-panel layout of Norton Commander.
Cross-platform development for Linux and macOS with support for
various architectures (ARM, x86).

## Core Features (MVP)

- **Two-Panel View:** Parallel display of two directories.
- **Navigation:** Quick directory switching via keyboard.
- **File Operations:** Copy (c), Move (r), Delete (d).
- **Platform Support:** Native behavior on Linux and macOS,
  support for ARM64 and x86_64 architectures.
- **Hidden Files:** Toggle visibility with 'h' key.

## Technology Stack

- **Language:** Go (Golang)
- **TUI Framework:** [Bubbletea](https://github.com/charmbracelet/bubbletea)
  (for the Elm Architecture Pattern)
- **Components:** [Bubbles](https://github.com/charmbracelet/bubbles)
  (for lists, inputs etc.)
- **Styling:** [Lip Gloss](https://github.com/charmbracelet/lipgloss)

## Implementation Plan

1. **Phase 1: Basic Framework**
   - Initialize Go project.
   - Set up Bubbletea basic loop.
   - Create layout for two panels (left/right).

2. **Phase 2: File System Navigation**
   - Function for reading directories.
   - Navigation (Up/Down, Enter to open, Backspace for back).
   - Focus switching between panels (Tab).

3. **Phase 3: File Operations**
   - Implementation of copy, move, and delete.
   - Progress display for large operations.

4. **Phase 4: Platform Optimization**
   - Linux support (ARM/x86).
   - macOS support.
   - Platform-specific paths and permissions.

5. **Phase 5: Extended Features**
   - Hidden file handling.
   - Toggle hidden files with 'h' key.

## Current Status

- **Completed:**
  - Basic framework with Bubbletea
  - Two-panel layout
  - Directory navigation
  - File operations (copy, move, delete)
  - Hidden file toggle
  - Cross-platform support (Linux, macOS)
  - Unit tests with 84.2% coverage (fs package)
  - Integration tests

- **Not Implemented:**
  - File viewing (removed)
  - Search functionality (removed)
  - Editor integration
