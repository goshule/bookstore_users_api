package services

import (
	"users_api/domain/users"
	"users_api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestError) {
	return user, nil
}
