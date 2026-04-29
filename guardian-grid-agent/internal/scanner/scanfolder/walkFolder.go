package scanfolder

import (
	"os"
	"path/filepath"
)

func walkFolder(root string, files *[]map[string]interface{}) error {
	maxFiles := 3000                   // prevent overload
	maxSize := int64(50 * 1024 * 1024) // 50MB max per file

	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if len(*files) >= maxFiles {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		fileData := map[string]interface{}{
			"name":       d.Name(),
			"path":       path,
			"size_bytes": info.Size(),
			"modified":   info.ModTime().Unix(),
		}

		// only hash if file size is reasonable
		if info.Size() <= maxSize {
			hash := hashFile(path)
			if hash != "" {
				fileData["sha256"] = hash
			}
		}

		*files = append(*files, fileData)

		return nil
	})
}
