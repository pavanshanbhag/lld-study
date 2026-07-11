package factory

import "fmt"

// PushNotification represents a push notification
type PushNotification struct{}

// NewPushNotification creates a new PushNotification instance
func NewPushNotification() *PushNotification {
	return &PushNotification{}
}

// Send sends a push notification
func (p *PushNotification) Send(message string) {
	fmt.Printf("Sending push notification: %s\n", message)
}
