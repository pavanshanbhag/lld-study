package pubsubsystem

import (
	"sync"
	"testing"
)

type recordingSubscriber struct {
	mu       sync.Mutex
	messages []*Message
}

func (s *recordingSubscriber) OnMessage(message *Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, message)
}

func (s *recordingSubscriber) count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.messages)
}

func TestTopicPublishToSubscribers(t *testing.T) {
	t.Parallel()

	topic := NewTopic("news")
	sub1 := &recordingSubscriber{}
	sub2 := &recordingSubscriber{}
	topic.AddSubscriber(sub1)
	topic.AddSubscriber(sub2)

	msg := NewMessage("hello")
	topic.Publish(msg)

	if sub1.count() != 1 || sub2.count() != 1 {
		t.Fatalf("expected 1 message each, got %d and %d", sub1.count(), sub2.count())
	}
}

func TestPublisherRegisteredTopic(t *testing.T) {
	t.Parallel()

	publisher := NewPublisher()
	topic := NewTopic("sports")
	publisher.RegisterTopic(topic)

	sub := &recordingSubscriber{}
	topic.AddSubscriber(sub)
	publisher.Publish(topic, NewMessage("score update"))

	if sub.count() != 1 {
		t.Fatalf("message count = %d, want 1", sub.count())
	}
}

func TestPublisherUnregisteredTopic(t *testing.T) {
	t.Parallel()

	publisher := NewPublisher()
	topic := NewTopic("weather")
	sub := &recordingSubscriber{}
	topic.AddSubscriber(sub)

	publisher.Publish(topic, NewMessage("sunny"))

	if sub.count() != 0 {
		t.Fatal("unregistered topic should not deliver messages")
	}
}

func TestTopicRemoveSubscriber(t *testing.T) {
	t.Parallel()

	topic := NewTopic("tech")
	sub := &recordingSubscriber{}
	topic.AddSubscriber(sub)
	topic.RemoveSubscriber(sub)
	topic.Publish(NewMessage("update"))

	if sub.count() != 0 {
		t.Fatal("removed subscriber should not receive messages")
	}
}
