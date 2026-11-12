package pubsubsystem

import "log/slog"

type PrintSubscriber struct {
	Name   string
	logger *slog.Logger
}

func NewPrintSubscriber(name string) *PrintSubscriber {
	return &PrintSubscriber{Name: name, logger: slog.Default()}
}

func (ps *PrintSubscriber) OnMessage(message *Message) {
	ps.logger.Info("message received", "subscriber", ps.Name, "content", message.Content)
}
