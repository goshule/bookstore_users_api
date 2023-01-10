package services

import (
	"log"
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
func UpdateUser(user *users.User, isPartial bool) (*users.User, *errors.RestError) {
	newUser := *user
	// log.Println("User-Service: ", newUser)
	_, err := GetUser(user)

	if err != nil {
		return nil, err
	}
	user.Id = newUser.Id
	if isPartial {
		if len(newUser.Email) > 0 {
			user.Email = newUser.Email
		}
		if len(newUser.FirstName) > 0 {
			user.Email = newUser.FirstName
		}
		if len(newUser.LastName) > 0 {
			user.Email = newUser.LastName
		}

	} else {

		user.Email = newUser.Email
		user.FirstName = newUser.FirstName
		user.LastName = newUser.LastName
	}

	log.Println("User-Service-New-user: ", newUser)

	errr := newUser.Update()

	if errr != nil {
		return nil, errr
	}

	return user, nil
}
