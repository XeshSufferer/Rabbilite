package rabbitmq

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	client *Client
}

func NewProducer(url string) (*Producer, error) {
	client, err := New(url)
	if err != nil {
		return nil, err
	}
	return &Producer{client: client}, nil
}

func (p *Producer) SendMessage(queueName string, message interface{}) error {

	_, err := p.client.channel.QueueDeclare(
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

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.client.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
			Timestamp:    time.Now(),
		})
	if err != nil {
		return err
	}

	log.Printf("Sent message to %s: %s", queueName, string(body))
	return nil
}

func (p *Producer) Close() {
	p.client.Close()
}
