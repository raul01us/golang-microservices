package domain

import (
	"fmt"
	"golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Raul", LastName: "Brindus", Email: "raul@xanda.net"},
	}
)

// GetUser retrieve from the DB the user identified by userID
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	// because of *User map type if wo dont find a user we will have "nil"
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
