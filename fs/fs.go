package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type FileEntry struct {
	Name  string
	IsDir bool
	Size  int64
}

func ReadDir(path string) ([]FileEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []FileEntry
	for _, entry := range entries {
		info, _ := entry.Info()
		files = append(files, FileEntry{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Size:  info.Size(),
		})
	}

	// Sort: directories first, then alphabetically
	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir != files[j].IsDir {
			return files[i].IsDir
		}
		return files[i].Name < files[j].Name
	})

	return files, nil
}

// Copy copies a file from src to dst
func Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// Move moves a file or directory from src to dst
func Move(src, dst string) error {
	// Try to rename the file (only works on the same partition)
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	// If rename fails (e.g. different partitions), copy and delete
	if err := Copy(src, dst); err != nil {
		return err
	}
	return os.Remove(src)
}

// Delete deletes a file or directory
func Delete(path string) error {
	return os.Remove(path)
}

// CopyDir copies a directory recursively
func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("src is not a directory: %s", src)
	}

	// Create destination directory if it doesn't exist
	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// Check if destination file already exists
			if _, err := os.Stat(dstPath); err == nil {
				return fmt.Errorf("destination file already exists: %s", dstPath)
			}
			err = Copy(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteDir deletes a directory recursively
func DeleteDir(path string) error {
	return os.RemoveAll(path)
}

// ReadDirRecursive reads a directory recursively (for future use)
func ReadDirRecursive(path string, depth int) ([]FileEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []FileEntry
	for _, entry := range entries {
		info, _ := entry.Info()
		isDir := entry.IsDir()
		files = append(files, FileEntry{
			Name:  entry.Name(),
			IsDir: isDir,
			Size:  info.Size(),
		})

		// Recursive for directories (up to a certain depth)
		if isDir && depth > 0 {
			subEntries, err := ReadDirRecursive(filepath.Join(path, entry.Name()), depth-1)
			if err != nil {
				continue
			}
			files = append(files, subEntries...)
		}
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir != files[j].IsDir {
			return files[i].IsDir
		}
		return files[i].Name < files[j].Name
	})

	return files, nil
}
