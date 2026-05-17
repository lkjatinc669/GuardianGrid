//go:build linux

package persistence

import (
	"os"
	"path/filepath"
)

func scanPersistence() ([]PersistenceItem, error) {
	items := make([]PersistenceItem, 0)

	// 1. Check systemd services (simplified: just list files in common dirs)
	systemdDirs := []string{
		"/etc/systemd/system",
		"/lib/systemd/system",
	}

	for _, dir := range systemdDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".service" {
				items = append(items, PersistenceItem{
					Type:     "Systemd",
					Name:     file.Name(),
					Location: filepath.Join(dir, file.Name()),
				})
			}
		}
	}

	// 2. Check Cron tabs
	cronDirs := []string{
		"/etc/cron.d",
		"/etc/cron.daily",
		"/etc/cron.hourly",
		"/etc/cron.monthly",
		"/etc/cron.weekly",
	}

	for _, dir := range cronDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() {
				items = append(items, PersistenceItem{
					Type:     "Cron",
					Name:     file.Name(),
					Location: filepath.Join(dir, file.Name()),
				})
			}
		}
	}

	// 3. XDG Autostart
	xdgDirs := []string{
		"/etc/xdg/autostart",
		filepath.Join(os.Getenv("HOME"), ".config/autostart"),
	}

	for _, dir := range xdgDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".desktop" {
				items = append(items, PersistenceItem{
					Type:     "Autostart",
					Name:     file.Name(),
					Location: filepath.Join(dir, file.Name()),
				})
			}
		}
	}

	return items, nil
}
