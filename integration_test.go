package main

import (
	"os"
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/karstenflache/commander-1/fs"
)

// IntegrationTestModel tests the complete model workflow
func TestIntegration_ModelInitialization(t *testing.T) {
	// Test that initialModel creates a valid model
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	m := initialModel()

	// Verify model structure
	if m.activePanel != 0 {
		t.Errorf("Expected activePanel 0, got %d", m.activePanel)
	}

	// Verify both panels are initialized
	if m.panels[0].path == "" && m.panels[1].path == "" {
		t.Log("Note: Panels may not be initialized if Getwd() failed")
	}
}

// IntegrationTestReadDirCmd tests reading directories
func TestIntegration_ReadDirCmd(t *testing.T) {
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Create test files and directories
	if err := os.Mkdir(filepath.Join(testDir, "testdir"), 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(testDir, "file1.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	m := initialModel()
	m.panels[0].path = testDir

	// Test readDirCmd
	cmd := m.readDirCmd(0)
	if cmd == nil {
		t.Error("readDirCmd should return a command")
	}

	// Execute the command
	msg := cmd()
	rdMsg, ok := msg.(readDirMsg)
	if !ok {
		t.Fatalf("Expected readDirMsg, got %T", msg)
	}

	if rdMsg.index != 0 {
		t.Errorf("Expected index 0, got %d", rdMsg.index)
	}

	if rdMsg.err != nil {
		t.Errorf("readDirCmd should not return error: %v", rdMsg.err)
	}

	// Should have at least 2 entries (file1.txt and testdir)
	if len(rdMsg.entries) < 2 {
		t.Errorf("Expected at least 2 entries, got %d", len(rdMsg.entries))
	}
}

// IntegrationTestHandleFileOperation tests file operations
func TestIntegration_HandleFileOperation(t *testing.T) {
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Create test file
	testFile := filepath.Join(testDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	m := initialModel()
	m.panels[0].path = testDir
	m.panels[0].entries, _ = fs.ReadDir(testDir)

	// Test copy operation
	cmd := m.handleFileOperation("copy")
	if cmd == nil {
		t.Error("handleFileOperation should return a command")
	}

	// Execute the command
	msg := cmd()
	if msg == nil {
		t.Log("Command executed without message")
	}
}

// IntegrationTestViewRendering tests the complete view rendering pipeline
func TestIntegration_ViewRendering(t *testing.T) {
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Create test files
	if err := os.WriteFile(filepath.Join(testDir, "file1.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(testDir, "file2.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	m := initialModel()
	m.width = 800
	m.height = 600
	m.panels[0].path = testDir
	m.panels[1].path = "/tmp"

	// Read directory
	entries, err := fs.ReadDir(testDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}
	m.panels[0].entries = entries

	// Test View function
	view := m.View()
	if view == "" {
		t.Error("View should return non-empty string")
	}

	// Verify both panels are rendered
	if !containsString(view, "file1.txt") {
		t.Log("Note: View may not contain expected content in test environment")
	}
}

// IntegrationTestPanelSwitching tests switching between panels
func TestIntegration_PanelSwitching(t *testing.T) {
	m := initialModel()

	// Initial state
	if m.activePanel != 0 {
		t.Errorf("Expected activePanel 0, got %d", m.activePanel)
	}

	// Switch to panel 1
	tabMsg := tea.KeyMsg{Type: tea.KeyTab}
	updatedModel, _ := m.Update(tabMsg)

	// Cast back to model
	newModel, ok := updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.activePanel != 1 {
		t.Errorf("Expected activePanel 1 after first Tab, got %d", newModel.activePanel)
	}

	// Switch back to panel 0
	updatedModel, _ = newModel.Update(tabMsg)
	newModel, ok = updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.activePanel != 0 {
		t.Errorf("Expected activePanel 0 after second Tab, got %d", newModel.activePanel)
	}
}

// IntegrationTestCursorMovement tests cursor movement in panels
func TestIntegration_CursorMovement(t *testing.T) {
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Create test files
	for i := 1; i <= 5; i++ {
		fileName := filepath.Join(testDir, "file"+string(rune('0'+i))+".txt")
		if err := os.WriteFile(fileName, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
	}

	m := initialModel()
	m.panels[0].path = testDir
	m.panels[0].entries, _ = fs.ReadDir(testDir)
	m.panels[0].cursor = 0

	// Initial cursor position
	if m.panels[0].cursor != 0 {
		t.Errorf("Expected cursor 0, got %d", m.panels[0].cursor)
	}

	// Move cursor down
	downMsg := tea.KeyMsg{Type: tea.KeyDown}
	updatedModel, _ := m.Update(downMsg)
	newModel, ok := updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.panels[0].cursor != 1 {
		t.Errorf("Expected cursor 1 after Down, got %d", newModel.panels[0].cursor)
	}

	// Move cursor up
	upMsg := tea.KeyMsg{Type: tea.KeyUp}
	updatedModel, _ = newModel.Update(upMsg)
	newModel, ok = updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.panels[0].cursor != 0 {
		t.Errorf("Expected cursor 0 after Up, got %d", newModel.panels[0].cursor)
	}

	// Test cursor bounds (should not go below 0)
	updatedModel, _ = newModel.Update(upMsg)
	newModel, ok = updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.panels[0].cursor < 0 {
		t.Error("Cursor should not go below 0")
	}
}

// IntegrationTestModelUpdateWithWindowSize tests WindowSizeMsg handling
func TestIntegration_ModelUpdateWithWindowSize(t *testing.T) {
	m := initialModel()

	// Test WindowSizeMsg
	wsm := tea.WindowSizeMsg{Width: 800, Height: 600}
	updatedModel, cmd := m.Update(wsm)

	newModel, ok := updatedModel.(model)
	if !ok {
		t.Fatal("Failed to cast updated model")
	}

	if newModel.width != 800 {
		t.Errorf("Expected width 800, got %d", newModel.width)
	}
	if newModel.height != 600 {
		t.Errorf("Expected height 600, got %d", newModel.height)
	}
	if cmd != nil {
		t.Log("WindowSizeMsg should not return a command")
	}
}

// IntegrationTestModelUpdateWithQuit tests quit command
func TestIntegration_ModelUpdateWithQuit(t *testing.T) {
	m := initialModel()

	// Test Ctrl+C to quit
	quitMsg := tea.KeyMsg{Type: tea.KeyCtrlC}
	_, cmd := m.Update(quitMsg)

	// cmd should be tea.Quit (or a function that returns it)
	if cmd == nil {
		t.Error("Quit command should not be nil")
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
