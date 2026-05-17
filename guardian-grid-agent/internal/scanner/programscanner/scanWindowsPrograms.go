//go:build windows

package programscanner

import (
	"bytes"
	"os/exec"
)

func scanPrograms() (map[string]interface{}, error) {
	ps := `
$paths = @(
  "HKLM:\Software\Microsoft\Windows\CurrentVersion\Uninstall\*",
  "HKLM:\Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\*",
  "HKCU:\Software\Microsoft\Windows\CurrentVersion\Uninstall\*"
)

$programs = foreach ($p in $paths) {
  Get-ItemProperty $p -ErrorAction SilentlyContinue |
  Where-Object { $_.DisplayName } |
  Select-Object DisplayName, DisplayVersion, Publisher
}

$programs | ConvertTo-Json -Compress
`
	cmd := exec.Command("powershell", "-Command", ps)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return emptyPrograms(), nil
	}

	return parseProgramsJSON(out.Bytes())
}
