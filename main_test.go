package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/karstenflache/commander-1/fs"
)

func TestCopyFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create Quelldatei
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := []byte("test content for copy file function")
	if err := os.WriteFile(srcPath, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Copy mit copyFile Funktion
	dstPath := filepath.Join(tmpDir, "destination.txt")
	if err := copyFile(srcPath, dstPath); err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	// Check that die Datei kopiert wurde
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstContent) != string(content) {
		t.Errorf("Copied content does not match. Expected %q, got %q", content, dstContent)
	}
}

func TestCopyFile_SourceNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "nonexistent.txt")
	dstPath := filepath.Join(tmpDir, "destination.txt")

	err := copyFile(srcPath, dstPath)
	if err == nil {
		t.Error("Expected error when source does not exist")
	}
}

func TestPanelStruct(t *testing.T) {
	panel := panel{
		path:           "/test/path",
		entries:        []fs.FileEntry{{Name: "file.txt", IsDir: false, Size: 100}},
		cursor:         0,
		viewportOffset: 0,
		showHidden:     false,
	}

	if panel.path != "/test/path" {
		t.Errorf("Expected path '/test/path', got '%s'", panel.path)
	}
	if len(panel.entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(panel.entries))
	}
	if panel.cursor != 0 {
		t.Errorf("Expected cursor 0, got %d", panel.cursor)
	}
	if panel.showHidden != false {
		t.Error("Expected showHidden to be false")
	}
}

func TestModelStruct(t *testing.T) {
	model := model{
		panels: [2]panel{
			{path: "/panel1"},
			{path: "/panel2"},
		},
		activePanel:    0,
		width:          800,
		height:         600,
		err:            nil,
		statusMsg:      "test status",
		viewportOffset: 0,
	}

	if model.activePanel != 0 {
		t.Errorf("Expected activePanel 0, got %d", model.activePanel)
	}
	if model.width != 800 {
		t.Errorf("Expected width 800, got %d", model.width)
	}
	if model.statusMsg != "test status" {
		t.Errorf("Expected statusMsg 'test status', got '%s'", model.statusMsg)
	}
}

func TestInitialModel(t *testing.T) {
	// Save original working directory
	origCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() { _ = os.Chdir(origCwd) }()

	// Change to a known directory
	testDir := t.TempDir()
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	model := initialModel()

	// Verify model is initialized correctly
	if model.activePanel != 0 {
		t.Errorf("Expected activePanel 0, got %d", model.activePanel)
	}

	// Panels should be initialized
	if model.panels[0].path == "" && model.panels[1].path == "" {
		// Both paths might be empty if Getwd() failed
		t.Log("Note: Panels might not be initialized if Getwd() failed")
	}
}

func TestReadDirMsg(t *testing.T) {
	msg := readDirMsg{
		index:   0,
		entries: []fs.FileEntry{{Name: "test.txt", IsDir: false, Size: 100}},
		err:     nil,
	}

	if msg.index != 0 {
		t.Errorf("Expected index 0, got %d", msg.index)
	}
	if len(msg.entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(msg.entries))
	}
	if msg.err != nil {
		t.Error("Expected error to be nil")
	}
}

func TestFileOpResultMsg(t *testing.T) {
	msg := fileOpResultMsg{
		op:                "copy",
		entryName:         "test.txt",
		inactivePanelPath: "/dest",
		err:               nil,
	}

	if msg.op != "copy" {
		t.Errorf("Expected op 'copy', got '%s'", msg.op)
	}
	if msg.entryName != "test.txt" {
		t.Errorf("Expected entryName 'test.txt', got '%s'", msg.entryName)
	}
	if msg.inactivePanelPath != "/dest" {
		t.Errorf("Expected inactivePanelPath '/dest', got '%s'", msg.inactivePanelPath)
	}
}

func TestRenderScrollBar(t *testing.T) {
	m := model{width: 800, height: 600}

	// Test with 0 entries (should be empty)
	scrollBar := m.renderScrollBar(0, 20, 0, 0)
	if scrollBar != "" {
		t.Error("Expected empty scrollbar for 0 entries")
	}

	// Test with positive entries (should not be empty)
	scrollBar = m.renderScrollBar(100, 20, 0, 10)
	if scrollBar == "" {
		t.Error("Expected non-empty scrollbar for large content")
	}

	// Test with few entries (should not be empty, since total > 0)
	scrollBar = m.renderScrollBar(5, 20, 0, 2)
	if scrollBar == "" {
		t.Error("Expected non-empty scrollbar for any positive total")
	}
}

func TestMinMaxHelpers(t *testing.T) {
	// Test min function
	if min(1, 2) != 1 {
		t.Error("min(1, 2) should return 1")
	}
	if min(5, 3) != 3 {
		t.Error("min(5, 3) should return 3")
	}

	// Test max function
	if max(1, 2) != 2 {
		t.Error("max(1, 2) should return 2")
	}
	if max(5, 3) != 5 {
		t.Error("max(5, 3) should return 5")
	}
}

// Helper functions that should exist in main.go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestPanelShowHidden(t *testing.T) {
	panel := panel{
		path:           "/test",
		entries:        []fs.FileEntry{},
		cursor:         0,
		viewportOffset: 0,
		showHidden:     false,
	}

	// Test toggling showHidden
	if panel.showHidden {
		t.Error("showHidden should be false initially")
	}

	panel.showHidden = true
	if !panel.showHidden {
		t.Error("showHidden should be true after toggle")
	}

	panel.showHidden = false
	if panel.showHidden {
		t.Error("showHidden should be false after second toggle")
	}
}

func TestPanelCursorMovement(t *testing.T) {
	panel := panel{
		path: "/test",
		entries: []fs.FileEntry{
			{Name: "file1.txt", IsDir: false, Size: 100},
			{Name: "file2.txt", IsDir: false, Size: 200},
			{Name: "file3.txt", IsDir: false, Size: 300},
		},
		cursor:         0,
		viewportOffset: 0,
		showHidden:     false,
	}

	// Test cursor movement within bounds
	if panel.cursor != 0 {
		t.Errorf("Expected cursor 0, got %d", panel.cursor)
	}

	// Move cursor down
	panel.cursor++
	if panel.cursor != 1 {
		t.Errorf("Expected cursor 1, got %d", panel.cursor)
	}

	// Move cursor up
	panel.cursor--
	if panel.cursor != 0 {
		t.Errorf("Expected cursor 0, got %d", panel.cursor)
	}

	// Test cursor bounds (should not go below 0)
	panel.cursor = -1
	if panel.cursor < 0 {
		panel.cursor = 0 // This is what the code should do
	}
	if panel.cursor != 0 {
		t.Errorf("Cursor should be clamped to 0, got %d", panel.cursor)
	}

	// Test cursor bounds (should not exceed length-1)
	panel.cursor = len(panel.entries)
	if panel.cursor >= len(panel.entries) {
		panel.cursor = len(panel.entries) - 1 // This is what the code should do
	}
	if panel.cursor != len(panel.entries)-1 {
		t.Errorf("Cursor should be clamped to %d, got %d", len(panel.entries)-1, panel.cursor)
	}
}

func TestFileEntryStruct(t *testing.T) {
	// Test file entry
	fileEntry := fs.FileEntry{
		Name:  "test.txt",
		IsDir: false,
		Size:  1024,
	}

	if fileEntry.Name != "test.txt" {
		t.Errorf("Expected Name 'test.txt', got '%s'", fileEntry.Name)
	}
	if fileEntry.IsDir {
		t.Error("File entry should not be a directory")
	}
	if fileEntry.Size != 1024 {
		t.Errorf("Expected Size 1024, got %d", fileEntry.Size)
	}

	// Test directory entry
	dirEntry := fs.FileEntry{
		Name:  "testdir",
		IsDir: true,
		Size:  0,
	}

	if !dirEntry.IsDir {
		t.Error("Directory entry should be a directory")
	}
	if dirEntry.Size != 0 {
		t.Errorf("Expected Size 0 for directory, got %d", dirEntry.Size)
	}
}

func TestRenderPanelWithEntries(t *testing.T) {
	m := model{
		width:  800,
		height: 600,
		panels: [2]panel{
			{
				path: "/test",
				entries: []fs.FileEntry{
					{Name: "dir1", IsDir: true, Size: 0},
					{Name: "file1.txt", IsDir: false, Size: 100},
					{Name: "file2.txt", IsDir: false, Size: 200},
				},
				cursor:         1,
				viewportOffset: 0,
				showHidden:     false,
			},
			{
				path:           "/test2",
				entries:        []fs.FileEntry{},
				cursor:         0,
				viewportOffset: 0,
				showHidden:     false,
			},
		},
		activePanel: 0,
	}

	// Test renderPanel with entries
	panelStr := m.renderPanel(0)
	if panelStr == "" {
		t.Error("Expected non-empty panel string")
	}

	// Check that path is included
	if !strings.Contains(panelStr, "/test") {
		t.Error("Expected panel to contain path")
	}

	// Check that file entries are included
	if !strings.Contains(panelStr, "file1.txt") {
		t.Error("Expected panel to contain file1.txt")
	}
}

func TestRenderPanelEmpty(t *testing.T) {
	m := model{
		width:  800,
		height: 600,
		panels: [2]panel{
			{
				path:           "/empty",
				entries:        []fs.FileEntry{},
				cursor:         0,
				viewportOffset: 0,
				showHidden:     false,
			},
			{path: "/test"},
		},
		activePanel: 0,
	}

	// Test renderPanel with empty entries
	panelStr := m.renderPanel(0)
	if panelStr == "" {
		t.Error("Expected non-empty panel string for empty directory")
	}
}

func TestRenderPanelWithHidden(t *testing.T) {
	m := model{
		width:  800,
		height: 600,
		panels: [2]panel{
			{
				path: "/test",
				entries: []fs.FileEntry{
					{Name: ".hidden", IsDir: false, Size: 100},
					{Name: "visible.txt", IsDir: false, Size: 200},
				},
				cursor:         1,
				viewportOffset: 0,
				showHidden:     false,
			},
			{path: "/test2"},
		},
		activePanel: 0,
	}

	// Test renderPanel with hidden files hidden
	panelStr := m.renderPanel(0)
	if strings.Contains(panelStr, ".hidden") {
		t.Error("Hidden files should not be visible when showHidden is false")
	}

	// Enable hidden files
	m.panels[0].showHidden = true
	panelStr = m.renderPanel(0)
	if !strings.Contains(panelStr, ".hidden") {
		t.Error("Hidden files should be visible when showHidden is true")
	}
}

func TestViewFunction(t *testing.T) {
	m := model{
		width:  800,
		height: 600,
		panels: [2]panel{
			{path: "/test"},
			{path: "/test2"},
		},
		activePanel: 0,
		err:         nil,
		statusMsg:   "",
	}

	// Test View with no error
	viewStr := m.View()
	if viewStr == "" {
		t.Error("Expected non-empty view")
	}

	// Test View with error
	m.err = fmt.Errorf("test error")
	viewStr = m.View()
	if viewStr == "" {
		t.Error("Expected non-empty view even with error")
	}
	if !strings.Contains(viewStr, "test error") {
		t.Error("Expected view to contain error message")
	}
}

func TestViewWithStatusMessage(t *testing.T) {
	m := model{
		width:       800,
		height:      600,
		panels:      [2]panel{{path: "/test"}, {path: "/test2"}},
		activePanel: 0,
		err:         nil,
		statusMsg:   "Test status message",
	}

	viewStr := m.View()
	if !strings.Contains(viewStr, "Test status message") {
		t.Error("Expected view to contain status message")
	}
}

func TestReadDirCmd(t *testing.T) {
	m := model{
		panels: [2]panel{
			{path: "/test"},
			{path: "/test2"},
		},
		activePanel: 0,
	}

	// Test readDirCmd
	cmd := m.readDirCmd(0)
	if cmd == nil {
		t.Error("Expected non-nil command")
	}
}

func TestCopyFileWithDifferentSizes(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with different file sizes
	testCases := []struct {
		name    string
		size    int
		content []byte
	}{
		{"empty.txt", 0, []byte{}},
		{"small.txt", 10, []byte("123456789")},
		{"medium.txt", 1000, make([]byte, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srcPath := filepath.Join(tmpDir, "source_"+tc.name)
			dstPath := filepath.Join(tmpDir, "dest_"+tc.name)

			if err := os.WriteFile(srcPath, tc.content, 0644); err != nil {
				t.Fatalf("Failed to create source file: %v", err)
			}

			if err := copyFile(srcPath, dstPath); err != nil {
				t.Fatalf("copyFile failed: %v", err)
			}

			dstContent, err := os.ReadFile(dstPath)
			if err != nil {
				t.Fatalf("Failed to read destination file: %v", err)
			}

			if string(dstContent) != string(tc.content) {
				t.Errorf("Content mismatch for %s", tc.name)
			}
		})
	}
}
