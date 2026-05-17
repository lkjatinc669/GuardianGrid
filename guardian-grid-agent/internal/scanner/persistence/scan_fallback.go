//go:build !windows && !linux && !darwin

package persistence

func scanPersistence() ([]PersistenceItem, error) {
	return []PersistenceItem{}, nil
}
