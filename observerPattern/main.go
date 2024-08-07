package main

import (
	"fmt"
	"learning/go/observerPattern/publisher"
	"learning/go/observerPattern/subscriber"
)

func main() {
	fmt.Println("Hello Observer!!")

	pub := publisher.NewPublisher()

	sub1 := subscriber.NewSubscriber("123")
	pub.AddSubscribers(sub1)

	sub2 := subscriber.NewSubscriber("456")
	pub.AddSubscribers(sub2)

	pub.BroadCastMessage("Hello Coders!!!")

	pub.RemoveSubscriber("123")

	pub.BroadCastMessage("Bye Coders!!!")

	fmt.Println("Bye Observer!!")
}
