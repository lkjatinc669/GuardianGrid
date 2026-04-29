package notification

import (
	"fmt"
	"runtime"
)

type NotificationType string

const (
	Info    NotificationType = "info"
	Warning NotificationType = "warning"
	Error   NotificationType = "error"
)

func ShowNotification(title, message string, nType NotificationType) {
	switch runtime.GOOS {

	case "windows":
		showWindows(title, message, nType)

	case "linux":
		showLinux(title, message, nType)

	case "darwin":
		showMac(title, message)

	default:
		fmt.Println("Unsupported OS")
	}
}
