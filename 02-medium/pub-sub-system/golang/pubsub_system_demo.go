package pubsubsystem

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	topic1 := NewTopic("Topic1")
	topic2 := NewTopic("Topic2")

	publisher1 := NewPublisher()
	publisher2 := NewPublisher()

	subscriber1 := NewPrintSubscriber("Subscriber1")
	subscriber2 := NewPrintSubscriber("Subscriber2")
	subscriber3 := NewPrintSubscriber("Subscriber3")

	publisher1.RegisterTopic(topic1)
	publisher2.RegisterTopic(topic2)

	topic1.AddSubscriber(subscriber1)
	topic1.AddSubscriber(subscriber2)
	topic2.AddSubscriber(subscriber2)
	topic2.AddSubscriber(subscriber3)

	if err := publisher1.Publish(topic1, NewMessage("Message1 for Topic1")); err != nil {
		logger.Error("publish failed", "err", err)
	}
	if err := publisher1.Publish(topic1, NewMessage("Message2 for Topic1")); err != nil {
		logger.Error("publish failed", "err", err)
	}
	if err := publisher2.Publish(topic2, NewMessage("Message1 for Topic2")); err != nil {
		logger.Error("publish failed", "err", err)
	}

	topic1.RemoveSubscriber(subscriber2)
	logger.Info("unsubscribed", "subscriber", subscriber2.Name, "topic", topic1.Name)

	if err := publisher1.Publish(topic1, NewMessage("Message3 for Topic1")); err != nil {
		logger.Error("publish failed", "err", err)
	}
	if err := publisher2.Publish(topic2, NewMessage("Message2 for Topic2")); err != nil {
		logger.Error("publish failed", "err", err)
	}
}
