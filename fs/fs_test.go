package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	// Test mit einem temporären Verzeichnis
	tmpDir := t.TempDir()

	// Erstelle einige Dateien und Verzeichnisse
	testData := []struct {
		name  string
		isDir bool
	}{
		{"file1.txt", false},
		{"file2.txt", false},
		{"subdir", true},
		{"a_file", false},
	}

	for _, data := range testData {
		path := filepath.Join(tmpDir, data.name)
		if data.isDir {
			if err := os.Mkdir(path, 0755); err != nil {
				t.Fatalf("Failed to create directory: %v", err)
			}
		} else {
			if err := os.WriteFile(path, []byte("test content"), 0644); err != nil {
				t.Fatalf("Failed to create file: %v", err)
			}
		}
	}

	// Test ReadDir
	entries, err := ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}

	if len(entries) != len(testData) {
		t.Errorf("Expected %d entries, got %d", len(testData), len(entries))
	}

	// Prüfe, dass Verzeichnisse zuerst kommen
	if len(entries) >= 2 {
		if entries[0].IsDir && !entries[1].IsDir {
			// OK - Verzeichnis kommt zuerst
		} else if !entries[0].IsDir && entries[1].IsDir {
			t.Error("Expected directories to come before files")
		}
	}

	// Prüfe, dass Dateien alphabetisch sortiert sind
	var fileNames []string
	for _, entry := range entries {
		if !entry.IsDir {
			fileNames = append(fileNames, entry.Name)
		}
	}
	for i := 1; i < len(fileNames); i++ {
		if fileNames[i] < fileNames[i-1] {
			t.Error("Files are not sorted alphabetically")
		}
	}
}

func TestReadDir_NonExistent(t *testing.T) {
	_, err := ReadDir("/non/existent/path")
	if err == nil {
		t.Error("Expected error for non-existent path")
	}
}

func TestCopy(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Quelldatei
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := []byte("test content for copy")
	if err := os.WriteFile(srcPath, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Kopiere
	dstPath := filepath.Join(tmpDir, "destination.txt")
	if err := Copy(srcPath, dstPath); err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	// Prüfe, dass die Datei kopiert wurde
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstContent) != string(content) {
		t.Errorf("Copied content does not match. Expected %q, got %q", content, dstContent)
	}

	// Prüfe, dass die ursprüngliche Datei noch existiert
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		t.Error("Source file should still exist after copy")
	}
}

func TestCopy_SourceNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "nonexistent.txt")
	dstPath := filepath.Join(tmpDir, "destination.txt")

	err := Copy(srcPath, dstPath)
	if err == nil {
		t.Error("Expected error when source does not exist")
	}
}

func TestMove(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Quelldatei
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := []byte("test content for move")
	if err := os.WriteFile(srcPath, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Verschiebe
	dstPath := filepath.Join(tmpDir, "destination.txt")
	if err := Move(srcPath, dstPath); err != nil {
		t.Fatalf("Move failed: %v", err)
	}

	// Prüfe, dass die Datei am neuen Ort existiert
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstContent) != string(content) {
		t.Errorf("Moved content does not match. Expected %q, got %q", content, dstContent)
	}

	// Prüfe, dass die ursprüngliche Datei nicht mehr existiert
	if _, err := os.Stat(srcPath); !os.IsNotExist(err) {
		t.Error("Source file should not exist after move")
	}
}

func TestMove_Directory(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Quelldatei
	srcPath := filepath.Join(tmpDir, "source.txt")
	if err := os.WriteFile(srcPath, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Verschiebe in neues Verzeichnis
	dstPath := filepath.Join(tmpDir, "newdir", "source.txt")
	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		t.Fatalf("Failed to create destination directory: %v", err)
	}

	if err := Move(srcPath, dstPath); err != nil {
		t.Fatalf("Move failed: %v", err)
	}

	// Prüfe, dass die Datei am neuen Ort existiert
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Error("Destination file should exist after move")
	}
}

func TestDelete(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Datei
	filePath := filepath.Join(tmpDir, "to_delete.txt")
	if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Lösche
	if err := Delete(filePath); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Prüfe, dass die Datei gelöscht wurde
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Error("File should not exist after delete")
	}
}

func TestDelete_NonExistent(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nonexistent.txt")

	err := Delete(filePath)
	if err == nil {
		t.Error("Expected error when deleting non-existent file")
	}
}

func TestCopyDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Quelldatei
	srcDir := filepath.Join(tmpDir, "source")
	if err := os.Mkdir(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Erstelle Dateien und Unterverzeichnisse
	if err := os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("content1"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "file2.txt"), []byte("content2"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	subDir := filepath.Join(srcDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "subfile.txt"), []byte("subcontent"), 0644); err != nil {
		t.Fatalf("Failed to create subfile: %v", err)
	}

	// Kopiere Verzeichnis
	dstDir := filepath.Join(tmpDir, "destination")
	if err := CopyDir(srcDir, dstDir); err != nil {
		t.Fatalf("CopyDir failed: %v", err)
	}

	// Prüfe, dass das Zielverzeichnis existiert
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		t.Error("Destination directory should exist")
	}

	// Prüfe, dass alle Dateien kopiert wurden
	files := []string{
		filepath.Join(dstDir, "file1.txt"),
		filepath.Join(dstDir, "file2.txt"),
		filepath.Join(dstDir, "subdir", "subfile.txt"),
	}

	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("File %s should exist after CopyDir", f)
		}
	}
}

func TestCopyDir_NotADirectory(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle eine Datei statt eines Verzeichnisses
	filePath := filepath.Join(tmpDir, "notadir")
	if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	err := CopyDir(filePath, filepath.Join(tmpDir, "destination"))
	if err == nil {
		t.Error("Expected error when source is not a directory")
	}
}

func TestCopyDir_DestinationExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Quellverzeichnis
	srcDir := filepath.Join(tmpDir, "source")
	if err := os.Mkdir(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Erstelle Zielverzeichnis bereits
	dstDir := filepath.Join(tmpDir, "destination")
	if err := os.Mkdir(dstDir, 0755); err != nil {
		t.Fatalf("Failed to create destination directory: %v", err)
	}

	err := CopyDir(srcDir, dstDir)
	if err == nil {
		t.Error("Expected error when destination already exists")
	}
}

func TestDeleteDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Verzeichnis mit Inhalt
	dirPath := filepath.Join(tmpDir, "to_delete")
	if err := os.Mkdir(dirPath, 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dirPath, "file.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Lösche Verzeichnis rekursiv
	if err := DeleteDir(dirPath); err != nil {
		t.Fatalf("DeleteDir failed: %v", err)
	}

	// Prüfe, dass das Verzeichnis gelöscht wurde
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		t.Error("Directory should not exist after DeleteDir")
	}
}

func TestReadDirRecursive(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Verzeichnisstruktur
	if err := os.MkdirAll(filepath.Join(tmpDir, "dir1", "subdir1"), 0755); err != nil {
		t.Fatalf("Failed to create directory structure: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(tmpDir, "dir2"), 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if err := os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "dir1", "file2.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "dir1", "subdir1", "file3.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "dir2", "file4.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Test mit Tiefe 1 (nur ein Level rekursiv)
	entries, err := ReadDirRecursive(tmpDir, 1)
	if err != nil {
		t.Fatalf("ReadDirRecursive failed: %v", err)
	}

	// Erwartet: dir1, dir2, file1.txt (dir1 enthält subdir1, wird nicht gelesen bei depth=1)
	if len(entries) < 3 {
		t.Errorf("Expected at least 3 entries, got %d", len(entries))
	}

	// Test mit Tiefe 2 (zwei Level rekursiv)
	entries, err = ReadDirRecursive(tmpDir, 2)
	if err != nil {
		t.Fatalf("ReadDirRecursive failed: %v", err)
	}

	// Erwartet mehr Einträge bei depth=2
	if len(entries) < 4 {
		t.Errorf("Expected at least 4 entries with depth=2, got %d", len(entries))
	}
}

func TestReadDirRecursive_NoDepth(t *testing.T) {
	tmpDir := t.TempDir()

	// Erstelle Verzeichnisstruktur
	if err := os.MkdirAll(filepath.Join(tmpDir, "dir1", "subdir1"), 0755); err != nil {
		t.Fatalf("Failed to create directory structure: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Test mit Tiefe 0 (nur das aktuelle Verzeichnis)
	entries, err := ReadDirRecursive(tmpDir, 0)
	if err != nil {
		t.Fatalf("ReadDirRecursive failed: %v", err)
	}

	// Nur das aktuelle Verzeichnis
	found := false
	for _, entry := range entries {
		if entry.Name == "file1.txt" {
			found = true
		}
		// Bei depth=0 sollten Unterverzeichnisse nicht rekursiv gelesen werden
		// aber der Eintrag selbst sollte vorhanden sein (dir1 in diesem Fall)
		_ = entry // Use entry to avoid unused variable warning
	}

	if !found {
		t.Error("Expected to find file1.txt")
	}
}

func TestFileEntry(t *testing.T) {
	entry := FileEntry{
		Name:  "test.txt",
		IsDir: false,
		Size:  1234,
	}

	if entry.Name != "test.txt" {
		t.Errorf("Expected Name 'test.txt', got '%s'", entry.Name)
	}
	if entry.IsDir != false {
		t.Error("Expected IsDir to be false")
	}
	if entry.Size != 1234 {
		t.Errorf("Expected Size 1234, got %d", entry.Size)
	}
}
