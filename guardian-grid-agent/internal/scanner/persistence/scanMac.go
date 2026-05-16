//go:build darwin

package persistence

import (
	"os"
	"path/filepath"
)

func scanMacPersistence() ([]PersistenceItem, error) {
	items := make([]PersistenceItem, 0)

	// LaunchAgents and LaunchDaemons
	dirs := []string{
		"/Library/LaunchAgents",
		"/Library/LaunchDaemons",
		filepath.Join(os.Getenv("HOME"), "Library/LaunchAgents"),
	}

	for _, dir := range dirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".plist" {
				items = append(items, PersistenceItem{
					Type:     "LaunchAgent/Daemon",
					Name:     file.Name(),
					Location: filepath.Join(dir, file.Name()),
				})
			}
		}
	}

	return items, nil
}
