package broker

import (
	"fmt"
	"sync"
	"testing"

	"github.com/streadway/amqp"
)

type myChannel struct {
}

func (*myChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	fmt.Println("this is mychannel")
	return nil, nil
}

func (*myChannel) Confirm(noWait bool) error {
	return nil
}

func (*myChannel) NotifyPublish(confirm chan amqp.Confirmation) chan amqp.Confirmation {
	return nil
}

func (*myChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return nil
}

func TestDialer(t *testing.T) {
	b := AMQPBroker{}

	c := myChannel{}

	b.Channel = &c
	GetMessages(&b, "hej")
}

func TestConfirmOne(t *testing.T) {

	// Maybe we want to check log messages?
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan amqp.Confirmation)
	go func(c <-chan amqp.Confirmation) {
		confirmOne(c)
		wg.Done()
	}(c)
	c <- amqp.Confirmation{}

	wg.Wait()
}
