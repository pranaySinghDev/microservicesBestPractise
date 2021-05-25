package service

import (
	"github.com/google/uuid"
	"github.com/pranaySinghDev/microservicesBestPractise/user-service/model"
)

func CreateUser(user *model.User) error {
	user.ID = uuid.New().String()
	return nil
}
