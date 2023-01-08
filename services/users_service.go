package services

import (
	"users_api/domain/users"
	"users_api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestError) {

	err := user.Validate()
	if err != nil {
		return nil, err
	}
	err = user.Save()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(user *users.User) (*users.User, *errors.RestError) {

	err := user.Get()

	if err != nil {
		return nil, err
	}

	return user, nil
}
func UpdateUser(user *users.User) (*users.User, *errors.RestError) {

	err := user.Update()

	if err != nil {
		return nil, err
	}

	return user, nil
}
