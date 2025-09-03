# Rabbilite

[![Go Reference](https://pkg.go.dev/badge/github.com/XeshSufferer/Rabbilite.svg)](https://pkg.go.dev/github.com/XeshSufferer/Rabbilite)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight and efficient library for working with RabbitMQ in Go applications. Rabbilite provides a simple interface for sending and consuming messages with support for persistence and acknowledgments.

## Features

- 🚀 Simple and intuitive API
- 💾 Persistent message support (Persistent)
- ✅ Automatic message acknowledgment/rejection (Ack/Nack)
- 🔄 Message requeue on errors
- 🧵 Thread-safe message processing
- 📦 JSON serialization out of the box

## Installation

```bash
go get github.com/rabbitmq/amqp091-go
```

```bash
go get github.com/XeshSufferer/Rabbilite
```

## Quick Start

### Sending Messages

```go
package main

import (
    "log"
    rabbitmq "github.com/XeshSufferer/Rabbilite/rabbilite"
)

func main() {
    producer, err := rabbitmq.NewProducer("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatal(err)
    }
    defer producer.Close()

    // Send a simple message
    err = producer.SendMessage("my_queue", "Hello, World!")
    if err != nil {
        log.Fatal(err)
    }

    // Send a struct
    type MyMessage struct {
        Name string `json:"name"`
        Value int    `json:"value"`
    }
    
    msg := MyMessage{Name: "test", Value: 42}
    err = producer.SendMessage("my_queue", msg)
}
```

### Consuming Messages

```go
package main

import (
    "log"
    rabbitmq "github.com/XeshSufferer/Rabbilite/rabbilite"
)

func main() {
    consumer, err := rabbitmq.NewConsumer("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatal(err)
    }
    defer consumer.Close()

    err = consumer.StartConsuming("my_queue", func(message []byte) error {
        log.Printf("Received: %s", string(message))
        // Your processing logic here
        return nil // Return nil acknowledges the message (Ack)
        // return errors.New("error") // Return error requeues the message (Nack + requeue)
    })
    
    if err != nil {
        log.Fatal(err)
    }

    // Wait indefinitely for messages
    select {}
}
```

## API Reference

### Producer

#### `NewProducer(url string) (*Producer, error)`
Creates a new producer with connection to RabbitMQ.

#### `producer.SendMessage(queueName string, message interface{}) error`
Sends a message to the specified queue. Supports any data types (serialized to JSON).

#### `producer.Close()`
Closes the connection to RabbitMQ.

### Consumer

#### `NewConsumer(url string) (*Consumer, error)`
Creates a new consumer with connection to RabbitMQ.

#### `consumer.StartConsuming(queueName string, handler MessageHandler) error`
Starts consuming messages from the queue. Handler must return `error` or `nil`.

#### `consumer.Close()`
Closes the connection to RabbitMQ.

### `MessageHandler` Type
```go
type MessageHandler func(message []byte) error
```

## Usage Example

```go
// producer.go
producer.SendMessage("tasks", map[string]interface{}{
    "task_id": 123,
    "action":  "process",
    "data":    "...",
})

// consumer.go
consumer.StartConsuming("tasks", func(message []byte) error {
    var task map[string]interface{}
    if err := json.Unmarshal(message, &task); err != nil {
        return err // Message will be requeued
    }
    
    // Process the task
    fmt.Printf("Processing task: %v\n", task)
    
    return nil // Message acknowledged
})
```

## RabbitMQ Configuration

The library works with standard RabbitMQ settings. For production use ensure:

1. Virtual host and user exist
2. Appropriate permissions are granted
3. Queue persistence is configured (already enabled in the library)

## Development

### Testing
```bash
# Run RabbitMQ in Docker
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

# Run tests
go test ./...
```

### Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License. See [LICENSE](LICENSE) file for details.

## Support

If you have questions or suggestions, create an issue in the repository or contact on Telegram: [@byxesh](https://t.me/byxesh)

---

**RabbiLite** - Making RabbitMQ in Go simple and efficient! 🐇✨
