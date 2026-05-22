package cve

import (
	"strings"
)

type Vulnerability struct {
	CVEID       string
	Severity    string
	Description string
	Score       float64
}

type CVEService struct {
	// Local DB: map[programName]map[version]Vulnerability
	localDB map[string]map[string][]Vulnerability
}

func NewCVEService() *CVEService {
	s := &CVEService{
		localDB: make(map[string]map[string][]Vulnerability),
	}
	s.seedLocalDB()
	return s
}

func (s *CVEService) seedLocalDB() {
	// Example vulnerabilities for testing
	s.addVulnerability("Google Chrome", "100.0.4896.60", Vulnerability{
		CVEID:       "CVE-2022-1096",
		Severity:    "High",
		Description: "Type Confusion in V8 in Google Chrome prior to 100.0.4896.60 allowed a remote attacker to potentially exploit heap corruption via a crafted HTML page.",
		Score:       8.8,
	})

	s.addVulnerability("Mozilla Firefox", "97.0", Vulnerability{
		CVEID:       "CVE-2022-26485",
		Severity:    "Critical",
		Description: "Use-after-free in XSLT parameter processing.",
		Score:       9.8,
	})

	s.addVulnerability("OpenSSL", "1.1.1k", Vulnerability{
		CVEID:       "CVE-2021-3711",
		Severity:    "High",
		Description: "SM2 Decryption Buffer Overflow.",
		Score:       8.1,
	})
}

func (s *CVEService) addVulnerability(name, version string, v Vulnerability) {
	if s.localDB[name] == nil {
		s.localDB[name] = make(map[string][]Vulnerability)
	}
	s.localDB[name][version] = append(s.localDB[name][version], v)
}

func (s *CVEService) Scan(name, version string) []Vulnerability {
	name = strings.ToLower(name)
	// Simple matching logic
	for progName, versions := range s.localDB {
		pName := strings.ToLower(progName)
		if strings.Contains(name, pName) || strings.Contains(pName, name) {
			if vulns, ok := versions[version]; ok {
				return vulns
			}
		}
	}
	return nil
}
