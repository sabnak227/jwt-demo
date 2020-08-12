package handlers

import (
	"encoding/json"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	userModels "github.com/sabnak227/jwt-demo/user/user-service/models"
	amqpAdapter "github.com/sabnak227/jwt-demo/util/amqp"
	"github.com/streadway/amqp"
)

func subscribers() {
	o := amqpAdapter.TopicSubscriber("user_create", "user_create.#")
	// use manual ack here
	o.ConsumeOptions.SetAutoAck(false)
	o.GenerateQueue = true
	o.QueueName = "user_create.auth_create"
	err := amqpClient.Subscribe(*o, createUserMsgProcessor)
	if err != nil {
		panic("Could not subscribe to exchange user_create")
	}
}

func createUserMsgProcessor(msg amqp.Delivery) {
	logger.Infof("Create user message received, %s", msg.Body)
	var user userModels.CreateUserMsg
	err := json.Unmarshal(msg.Body, &user)
	if err != nil {
		logger.Errorf("Failed to unmarshal create auth message for user %s, requeueing...., error %s", user.Email, err)
		rejectMsg(msg)
		return
	}

	if err := repo.CreateAuth(models.Auth{
		UserID:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}); err != nil {
		logger.Errorf("Failed to create auth for user %s, requeueing...., error %s", user.Email, err)
		rejectMsg(msg)
		return
	}
	ackMsg(msg, true)
}

func ackMsg(msg amqp.Delivery, multiple bool) {
	if err := msg.Ack(multiple); err != nil {
		logger.Fatalf("Failed to acknowledge message {%s}, error %s", msg.Body, err)
	}
}
func rejectMsg(msg amqp.Delivery) {
	if err := msg.Nack(false, true); err != nil {
		logger.Fatalf("Failed to reject message {%s}, error %s", msg.Body, err)
	}
}