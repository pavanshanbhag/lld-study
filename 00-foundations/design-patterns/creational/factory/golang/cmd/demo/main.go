package main

import (
	"fmt"

	"factory"
)

func main() {
	notificationFactory := factory.NewSimpleNotificationFactory()

	notifications := []string{"EMAIL", "SMS", "PUSH"}
	messages := []string{
		"Welcome to our platform!",
		"Your OTP is 123456",
		"You have a new follower!",
	}

	for i, notificationType := range notifications {
		notification, err := notificationFactory.CreateNotification(notificationType)
		if err != nil {
			fmt.Printf("Error creating notification: %v\n", err)
			continue
		}
		notification.Send(messages[i])
	}
}
