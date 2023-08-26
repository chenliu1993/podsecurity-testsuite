package files

import (
	"os"
	"path/filepath"
)

// WalkPath returns a list of all files in the given root directory
// Caller should use absolute paths
func WalkPath(root string) ([]string, error) {
	info, err := os.Lstat(root)
	if err != nil {
		return []string{}, err
	}

	// only care about files that has the expected file names
	if !info.IsDir() {
		if (info.Mode()&os.ModeSymlink == os.ModeSymlink) ||
			((filepath.Ext(info.Name()) != ".yaml") &&
				(filepath.Ext(info.Name()) != ".yml")) {
			return []string{}, nil
		}

		return []string{root}, nil
	}

	// recursively walk through the directory
	dirs, err := os.ReadDir(root)
	if err != nil {
		return []string{}, err
	}

	var allFiles []string
	for _, dir := range dirs {
		files, err := WalkPath(filepath.Join(root, dir.Name()))
		if err != nil {
			return []string{}, err
		}
		allFiles = append(allFiles, files...)
	}
	return allFiles, nil
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, os.FileMode(0644))
}

// Caller should use this function to clean up all the resource files
func Cleanup(path string) error {
	return os.RemoveAll(path)
}

func CheckDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(path, os.FileMode(0755)); err != nil {
				return err
			}
		}
		return err
	}
	return nil
}
