package cve

import (
	"testing"
)

func TestCVEScan(t *testing.T) {
	s := NewCVEService()

	tests := []struct {
		name     string
		version  string
		expected int
	}{
		{"Google Chrome", "100.0.4896.60", 1},
		{"Mozilla Firefox", "97.0", 1},
		{"OpenSSL", "1.1.1k", 1},
		{"Some Safe App", "1.0.0", 0},
		{"Chrome", "100.0.4896.60", 1}, // partial match
	}

	for _, tt := range tests {
		findings := s.Scan(tt.name, tt.version)
		if len(findings) != tt.expected {
			t.Errorf("Scan(%s, %s) = %d findings, want %d", tt.name, tt.version, len(findings), tt.expected)
		}
	}
}
