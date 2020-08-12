package amqp

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

// IAmqpClient Defines our interface for connecting, producing and consuming messages.
type IAmqpClient interface {
	ConnectToBroker(connectionString string) error
	Publish(options PublisherOptions, body []byte, routingKey string) error
	Subscribe(options SubscriberOptions, handlerFunc func(amqp.Delivery)) error
	Close()
}

// AmqpClient  Real implementation, encapsulates a pointer to an amqp.Connection
type AmqpClient struct {
	conn *amqp.Connection
}

// ConnectToBroker connects to an AMQP broker using the supplied connectionString.
func (m *AmqpClient) ConnectToBroker(connectionString string) error {
	if connectionString == "" {
		return fmt.Errorf("empty dns")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))

	return err
}


type PublisherOptions struct {
	ExchangeName string
	ExchangeType string
	BindingKey   string
	GenerateQueue bool
	QueueName string
	QueueOptions *QueueOptions
	QueueBindOptions *QueueBindOptions
	ExchangeOptions *ExchangeOptions
	PublishOptions *PublishOptions
}

//TopicPublisher initiates a opic type exchange
func TopicPublisher(exchangeName string, bindingKey string) *PublisherOptions {
	var o PublisherOptions
	o.ExchangeName = exchangeName
	o.ExchangeType = amqp.ExchangeTopic
	o.BindingKey = bindingKey
	o.ExchangeOptions = &ExchangeOptions{
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	o.PublishOptions = &PublishOptions{
		Mandatory: false,
		Immediate: false,
	}
	o.GenerateQueue = false
	o.QueueName = ""
	o.QueueOptions = &QueueOptions{
		Durable: true,
		AutoDelete: false,
		Exclusive: false,
		NoWait: false,
		Args: nil,
	}
	o.QueueBindOptions = &QueueBindOptions{
		NoWait: false,
		Args: nil,
	}

	return &o
}


type SubscriberOptions struct {
	ExchangeName string
	ExchangeType string
	BindingKey   string
	GenerateQueue bool
	QueueName string
	QueueOptions *QueueOptions
	QueueBindOptions *QueueBindOptions
	ExchangeOptions *ExchangeOptions
	ConsumeOptions *ConsumeOptions
}

func TopicSubscriber(exchangeName string, bindingKey string) *SubscriberOptions {
	var o SubscriberOptions
	o.ExchangeName = exchangeName
	o.ExchangeType = amqp.ExchangeTopic
	o.BindingKey = bindingKey
	o.ExchangeOptions = &ExchangeOptions{
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	o.ConsumeOptions = &ConsumeOptions{
		ConsumerName: "",
		AutoAck: true,
		Exclusive: false,
		NoLocal: false,
		NoWait: false,
		Args: nil,
	}
	o.GenerateQueue = false
	o.QueueName = ""
	o.QueueOptions = &QueueOptions{
		Durable: true,
		AutoDelete: false,
		Exclusive: false,
		NoWait: false,
		Args: nil,
	}
	o.QueueBindOptions = &QueueBindOptions{
		NoWait: false,
		Args: nil,
	}

	return &o
}
// TODO: implement direct, fanout and header types as well

// Publish publishes a message to the named exchange.
func (m *AmqpClient) Publish(options PublisherOptions, body []byte, routingKey string) error {
	if m.conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to connect to channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		options.ExchangeName,
		options.ExchangeType,
		options.ExchangeOptions.Durable,
		options.ExchangeOptions.AutoDelete,
		options.ExchangeOptions.Internal,
		options.ExchangeOptions.NoWait,
		options.ExchangeOptions.Args,
	)
	failOnError(err, "Failed to register an Exchange")

	if options.GenerateQueue {
		log(fmt.Sprintf("declared Exchange, declaring Queue (%s)", ""))
		queue, err := ch.QueueDeclare(
			options.QueueName,
			options.QueueOptions.Durable,
			options.QueueOptions.AutoDelete,
			options.QueueOptions.Exclusive,
			options.QueueOptions.NoWait,
			options.QueueOptions.Args,
		)
		failOnError(err, "Failed to declare queue")
		options.QueueName = queue.Name

		if err := ch.QueueBind(
			options.QueueName,
			options.BindingKey,
			options.ExchangeName,
			options.QueueBindOptions.NoWait,
			options.QueueBindOptions.Args,
		); err != nil {
			failOnError(err, "Failed to bind queue")
		}
	}


	err = ch.Publish(
		options.ExchangeName,
		routingKey,
		options.PublishOptions.Mandatory,
		options.PublishOptions.Immediate,
		amqp.Publishing{
			Body: body,
		})
	log(fmt.Sprintf("A message was sent: %v", string(body)))
	return err
}

// Subscribe registers a handler function for a given exchange.
func (m *AmqpClient) Subscribe(options SubscriberOptions, handlerFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		options.ExchangeName,
		options.ExchangeType,
		options.ExchangeOptions.Durable,
		options.ExchangeOptions.AutoDelete,
		options.ExchangeOptions.Internal,
		options.ExchangeOptions.NoWait,
		options.ExchangeOptions.Args,
	)
	failOnError(err, "Failed to register an Exchange")

	if options.GenerateQueue {
		log(fmt.Sprintf("declared Exchange, declaring Queue (%s)", ""))
		queue, err := ch.QueueDeclare(
			options.QueueName,
			options.QueueOptions.Durable,
			options.QueueOptions.AutoDelete,
			options.QueueOptions.Exclusive,
			options.QueueOptions.NoWait,
			options.QueueOptions.Args,
		)
		failOnError(err, "Failed to declare queue")
		options.QueueName = queue.Name

		if err := ch.QueueBind(
			options.QueueName,
			options.BindingKey,
			options.ExchangeName,
			options.QueueBindOptions.NoWait,
			options.QueueBindOptions.Args,
		); err != nil {
			failOnError(err, "Failed to bind queue")
		}
	}

	msgs, err := ch.Consume(
		options.QueueName,
		options.ConsumeOptions.ConsumerName,
		options.ConsumeOptions.AutoAck,
		options.ConsumeOptions.Exclusive,
		options.ConsumeOptions.NoLocal,
		options.ConsumeOptions.NoWait,
		options.ConsumeOptions.Args,
	)
	failOnError(err, "Failed to register a consumer")

	go consumeLoop(msgs, ch, handlerFunc)
	return nil
}

// Close closes the connection to the AMQP-broker, if available.
func (m *AmqpClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, ch *amqp.Channel, handlerFunc func(d amqp.Delivery)) {
	defer ch.Close()
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
	logrus.Infof("AMQP Message: %s\n", msg)
}

type ExchangeOptions struct {
	Durable bool
	AutoDelete bool
	Internal bool
	NoWait bool
	Args amqp.Table
}

func (eo *ExchangeOptions) SetExchangeDurable(flag bool) {
	eo.Durable = flag
}

func (eo *ExchangeOptions) SetExchangeAutoDelete(flag bool) {
	eo.AutoDelete = flag
}

func (eo *ExchangeOptions) SetExchangeInternal(flag bool) {
	eo.Internal = flag
}

func (eo *ExchangeOptions) SetExchangeNoWait(flag bool) {
	eo.NoWait = flag
}

func (eo *ExchangeOptions) SetExchangeArgs(args amqp.Table) {
	eo.Args = args
}

type PublishOptions struct {
	Mandatory bool
	Immediate bool
}

func (po *PublishOptions) SetMandatory(flag bool) {
	po.Mandatory = flag
}

func (po *PublishOptions) SetImmediate(flag bool) {
	po.Immediate = flag
}

type ConsumeOptions struct {
	ConsumerName string
	AutoAck bool
	Exclusive bool
	NoLocal bool
	NoWait bool
	Args amqp.Table
}

func (co *ConsumeOptions) SetConsumerName(name string) {
	co.ConsumerName = name
}

func (co *ConsumeOptions) SetAutoAck(flag bool) {
	co.AutoAck = flag
}

func (co *ConsumeOptions) SetExclusive(flag bool) {
	co.Exclusive = flag
}

func (co *ConsumeOptions) SetNoLocal(flag bool) {
	co.NoLocal = flag
}

func (co *ConsumeOptions) SetNoWait(flag bool) {
	co.NoWait = flag
}

func (co *ConsumeOptions) SetArgs(args amqp.Table) {
	co.Args = args
}

type QueueOptions struct {
	Durable bool
	AutoDelete bool
	Exclusive bool
	NoWait bool
	Args amqp.Table
}

func (qo *QueueOptions) SetDurable(flag bool) {
	qo.Durable = flag
}

func (qo *QueueOptions) SetAutoDelete(flag bool) {
	qo.AutoDelete = flag
}

func (qo *QueueOptions) SetExclusive(flag bool) {
	qo.Exclusive = flag
}

func (qo *QueueOptions) SetNoWait(flag bool) {
	qo.NoWait = flag
}

func (qo *QueueOptions) SetArgs(args amqp.Table) {
	qo.Args = args
}

type QueueBindOptions struct {
	NoWait bool
	Args amqp.Table
}

func (qbo *QueueBindOptions) SetNoWait(flag bool) {
	qbo.NoWait = flag
}

func (qbo *QueueBindOptions) SetArgs(args amqp.Table) {
	qbo.Args = args
}