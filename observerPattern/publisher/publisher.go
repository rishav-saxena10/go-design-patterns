package publisher

import "learning/go/observerPattern/subscriber"

type PublisherInterface interface {
	AddSubscribers(sub subscriber.SubscriberInterface)
	RemoveSubscriber(subId string)
	BroadCastMessage(message string)
}

type Publisher struct {
	subscribers map[string]subscriber.SubscriberInterface
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string]subscriber.SubscriberInterface),
	}
}

func (p *Publisher) AddSubscribers(subscriber subscriber.SubscriberInterface) {
	p.subscribers[subscriber.Id()] = subscriber
}

func (p *Publisher) RemoveSubscriber(subId string) {
	delete(p.subscribers, subId)
}

func (p *Publisher) BroadCastMessage(message string) {
	for _, subscriber := range p.subscribers {
		subscriber.React(message)
	}
}
