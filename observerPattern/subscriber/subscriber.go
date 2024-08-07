package subscriber

import "fmt"

type SubscriberInterface interface {
	Id() string
	React(message string)
}

type Subscriber struct {
	id string
}

func NewSubscriber(subId string) *Subscriber {
	return &Subscriber{
		id: subId,
	}
}

func (s *Subscriber) Id() string {
	return s.id
}

func (s *Subscriber) React(message string) {
	fmt.Printf("Subscriber Id: %s, received message: %s", s.Id(), message)
	fmt.Println()
}
