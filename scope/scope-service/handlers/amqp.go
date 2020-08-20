package handlers

import (
	"encoding/json"
	"fmt"
	userModels "github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/streadway/amqp"
)

func subscribers() {
	o := amqpAdapter.FanoutSubscriber("user_updates")
	// use manual ack here
	o.ConsumeOptions.SetAutoAck(false)
	o.GenerateQueue = true
	o.QueueName = "user_updates.scope_updates"
	err := amqpClient.Subscribe(*o, UserUpdateMsgProcessor)
	if err != nil {
		panic("Could not subscribe to exchange user_create")
	}
}

func UserUpdateMsgProcessor(msg amqp.Delivery) {
	logger.Infof("Create user message received, %s", msg.Body)
	var user userModels.UserMsg
	if err := json.Unmarshal(msg.Body, &user); err != nil {
		logger.Errorf("Failed to unmarshal user svc message for user %s, error %s", msg.Body, err)
		amqpAdapter.NackMsg(msg, logger)
		return
	}

	var err error
	switch user.Type {
	case userModels.UserMsgTypeCreated:
		err = createDefaultScope(user)
	default:
		err = fmt.Errorf("undefined message type: %s", user.Type)
	}
	if err != nil {
		logger.Errorf("Failed to unmarshal user svc message for user %s, error %s", msg.Body, err)
		amqpAdapter.NackMsg(msg, logger)
		return
	}
	amqpAdapter.AckMsg(msg, logger)
}

func createDefaultScope(user userModels.UserMsg) error {
	return repo.AssignRole(repo.GetConn(), user.UserId, "user")
}