package amqp

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func AckMsg(msg amqp.Delivery, logger *logrus.Logger) {
	if err := msg.Ack(false); err != nil {
		logger.Errorf("Failed to acknowledge message {%s}, error %s", msg.Body, err)
	}
}

func NackMsg(msg amqp.Delivery, logger *logrus.Logger) {
	if err := msg.Nack(false, false); err != nil {
		logger.Errorf("Failed to reject message {%s}, error %s", msg.Body, err)
	}
}

func RejectMsg(msg amqp.Delivery, logger *logrus.Logger) {
	if err := msg.Reject(false); err != nil {
		logger.Errorf("Failed to reject message {%s}, error %s", msg.Body, err)
	}
}