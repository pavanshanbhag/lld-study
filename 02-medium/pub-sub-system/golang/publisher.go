package pubsubsystem

import "fmt"

type Publisher struct {
	Topics map[*Topic]struct{}
}

func NewPublisher() *Publisher {
	return &Publisher{Topics: make(map[*Topic]struct{})}
}

func (p *Publisher) RegisterTopic(topic *Topic) {
	p.Topics[topic] = struct{}{}
}

func (p *Publisher) Publish(topic *Topic, message *Message) error {
	if _, exists := p.Topics[topic]; !exists {
		return fmt.Errorf("publisher not registered for topic %s", topic.Name)
	}
	topic.Publish(message)
	return nil
}
