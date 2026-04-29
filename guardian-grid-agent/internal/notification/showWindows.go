package notification

import (
	"fmt"
	"os/exec"
)

func showWindows(title, message string, nType NotificationType) {
	prefix := ""

	switch nType {
	case Error:
		prefix = "❌ "
	case Warning:
		prefix = "⚠️ "
	default:
		prefix = "ℹ️ "
	}

	cmd := exec.Command("powershell",
		"-Command",
		fmt.Sprintf(`New-BurntToastNotification -Text "%s", "%s"`,
			title, prefix+message),
	)

	cmd.Run()
}
