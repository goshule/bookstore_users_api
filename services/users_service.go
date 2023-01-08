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

	us, err := GetUser(user)

	if err != nil {
		return nil, err
	}

	currentUser := users.User{}
	currentUser.Id = user.Id

	if len(us.Email) > 0 {
		currentUser.Email = us.Email
	}
	if len(us.FirstName) > 0 {
		currentUser.FirstName = us.FirstName
	}
	if len(us.LastName) > 0 {
		currentUser.LastName = us.LastName
	}

	errr := currentUser.Update()

	if errr != nil {
		return nil, err
	}

	return &currentUser, nil
}
