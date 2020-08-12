package amqp

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

// IAmqpClient Defines our interface for connecting, producing and consuming messages.
type IAmqpClient interface {
	ConnectToBroker(connectionString string)
	Publish(msg []byte, exchangeName string, exchangeType string) error
	PublishOnQueue(msg []byte, queueName string) error
	PublishOnQueueWithContext(ctx context.Context, msg []byte, queueName string) error
	Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error
	SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error
	Close()
}

// AmqpClient  Real implementation, encapsulates a pointer to an amqp.Connection
type AmqpClient struct {
	conn *amqp.Connection
}

// ConnectToBroker connects to an AMQP broker using the supplied connectionString.
func (m *AmqpClient) ConnectToBroker(connectionString string) {
	if connectionString == "" {
		panic("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		panic("Failed to connect to AMQP compatible broker at: " + connectionString)
	}
}

// Publish publishes a message to the named exchange.
func (m *AmqpClient) Publish(body []byte, exchangeName string, exchangeType string) error {
	if m.conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	defer ch.Close()
	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	failOnError(err, "Failed to register an Exchange")

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		"",    // our queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)

	err = ch.Publish( // Publishes a message onto the queue.
		exchangeName, // exchange
		exchangeName, // routing key      q.Name
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Body: body, // Our JSON body as []byte
		})
	log(fmt.Sprintf("A message was sent: %v", string(body)))
	return err
}

// PublishOnQueueWithContext publishes the supplied body onto the named queue, passing the context.
func (m *AmqpClient) PublishOnQueueWithContext(ctx context.Context, body []byte, queueName string) error {
	if m.conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	defer ch.Close()

	queue, err := ch.QueueDeclare( // Declare a queue that will be created if not exists with some args
		queueName, // our queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	// Publishes a message onto the queue.
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		buildMessage(ctx, body))
	log(fmt.Sprintf("A message was sent to queue %v: %v", queueName, string(body)))
	return err
}

func buildMessage(ctx context.Context, body []byte) amqp.Publishing {
	publishing := amqp.Publishing{
		ContentType: "application/json",
		Body:        body, // Our JSON body as []byte
	}
	// if ctx != nil {
	// 	child := tracing.StartChildSpanFromContext(ctx, "messaging")
	// 	defer child.Finish()
	// 	var val = make(opentracing.TextMapCarrier)
	// 	err := tracing.AddTracingToTextMapCarrier(child, val)
	// 	if err != nil {
	// 		logrus.Errorf("Error injecting span context: %v", err.Error())
	// 	} else {
	// 		publishing.Headers = tracing.CarrierToMap(val)
	// 	}
	// }
	return publishing
}

// PublishOnQueue publishes the supplied body on the queueName.
func (m *AmqpClient) PublishOnQueue(body []byte, queueName string) error {
	return m.PublishOnQueueWithContext(nil, body, queueName)
}

// Subscribe registers a handler function for a given exchange.
func (m *AmqpClient) Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	failOnError(err, "Failed to register an Exchange")

	log(fmt.Sprintf("declared Exchange, declaring Queue (%s)", ""))
	queue, err := ch.QueueDeclare(
		"",    // name of the queue
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to register an Queue")

	log(fmt.Sprintf("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
		queue.Messages, queue.Consumers, exchangeName))

	err = ch.QueueBind(
		queue.Name,   // name of the queue
		exchangeName, // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	go consumeLoop(msgs, handlerFunc)
	return nil
}

// SubscribeToQueue registers a handler function for the named queue.
func (m *AmqpClient) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")

	log(fmt.Sprintf("Declaring Queue (%s)", queueName))
	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	failOnError(err, "Failed to register an Queue")

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	go consumeLoop(msgs, handlerFunc)
	return nil
}

// Close closes the connection to the AMQP-broker, if available.
func (m *AmqpClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Errorf("AMQP Failure %s: %s", msg, err)
	}
}

func log(msg string) {
	logrus.Infof("msg:%s\n", msg)
}