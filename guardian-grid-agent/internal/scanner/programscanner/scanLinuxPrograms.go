package programscanner

func scanLinuxPrograms() (map[string]interface{}, error) {
	// Try dpkg (Debian/Ubuntu)
	if data, err := runCmd("dpkg-query", "-W", "-f=${Package}\t${Version}\n"); err == nil {
		return parseKeyValueLines(data)
	}

	// Try rpm (RHEL/Fedora)
	if data, err := runCmd("rpm", "-qa", "--qf", "%{NAME}\t%{VERSION}\n"); err == nil {
		return parseKeyValueLines(data)
	}

	return emptyPrograms(), nil
}
