package notification

type NotificationType string

const (
	Info    NotificationType = "info"
	Warning NotificationType = "warning"
	Error   NotificationType = "error"
)

func ShowNotification(title, message string, nType NotificationType) {
	showNotification(title, message, nType)
}
