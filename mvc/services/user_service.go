package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
)

// GetUser makes the call to the domain that handles GetUser
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
