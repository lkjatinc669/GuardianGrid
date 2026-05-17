//go:build windows

package persistence

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func scanPersistence() ([]PersistenceItem, error) {
	items := make([]PersistenceItem, 0)

	// 1. Registry keys for "Run"
	// ... (rest of registry logic)
	keys := []struct {
		root registry.Key
		path string
	}{
		{registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`},
		{registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce`},
		{registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`},
		{registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce`},
		{registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`},
	}

	for _, k := range keys {
		regKey, err := registry.OpenKey(k.root, k.path, registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		defer regKey.Close()

		names, err := regKey.ReadValueNames(-1)
		if err != nil {
			continue
		}

		for _, name := range names {
			val, _, err := regKey.GetStringValue(name)
			if err != nil {
				continue
			}

			items = append(items, PersistenceItem{
				Type:     "Registry",
				Name:     name,
				Path:     val,
				Location: k.path,
			})
		}
	}

	// 2. Startup Folders
	startupDirs := []string{
		filepath.Join(os.Getenv("ProgramData"), `Microsoft\Windows\Start Menu\Programs\StartUp`),
		filepath.Join(os.Getenv("AppData"), `Microsoft\Windows\Start Menu\Programs\Startup`),
	}

	for _, dir := range startupDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() {
				items = append(items, PersistenceItem{
					Type:     "StartupFolder",
					Name:     file.Name(),
					Location: filepath.Join(dir, file.Name()),
				})
			}
		}
	}

	return items, nil
}
