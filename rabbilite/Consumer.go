package rabbitmq

import (
	"log"
)

type Consumer struct {
	client *Client
}

func NewConsumer(url string) (*Consumer, error) {
	client, err := New(url)
	if err != nil {
		return nil, err
	}
	return &Consumer{client: client}, nil
}

type MessageHandler func(message []byte) error

func (c *Consumer) StartConsuming(queueName string, handler MessageHandler) error {

	_, err := c.client.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := c.client.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message from %s: %s", queueName, string(msg.Body))

			if err := handler(msg.Body); err != nil {
				log.Printf("Error handling message: %v", err)
				msg.Nack(false, true)
			} else {
				msg.Ack(false)
			}
		}
	}()

	log.Printf("Started consuming from queue: %s", queueName)
	return nil
}

func (c *Consumer) Close() {
	c.client.Close()
}
