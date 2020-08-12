package handlers

import (
	"encoding/json"
	"github.com/sabnak227/jwt-demo/auth/auth-service/models"
	userModels "github.com/sabnak227/jwt-demo/user/user-service/models"
	"github.com/streadway/amqp"
)

func subscribers() {
	// Call the subscribe method with queue name and callback function
	err := amqpClient.SubscribeToQueue("hello", "hello", createUserMsgProcessor)
	if err != nil {
		panic("Could not subscribe to queue updateFeedQueue")
	}
}

func createUserMsgProcessor(msg amqp.Delivery) {
	logger.Infof("Create user message received, %s", msg.Body)

	var user userModels.CreateUserMsg
	err := json.Unmarshal(msg.Body, &user)
	if err != nil {
		//todo requeue
	}

	if err := repo.CreateAuth(models.Auth{
		UserID:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}); err != nil {
		//todo requeue
	}
}