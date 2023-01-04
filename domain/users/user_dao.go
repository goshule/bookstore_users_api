package users

import (
	"fmt"
	"log"
	"users_api/datasources/mysql/users_db"
	"users_api/utils/date_utils"
	"users_api/utils/errors"
)

var userDB = make(map[int64]*User)

const (
	queryInsertUser = "insert into users(FirstName,LastName,email,DateCreated)values(?,?,?,?);"
	querySelectUser = "select id,FirstName,LastName,email,DateCreated from users where id=?"
)

func (user *User) Save() *errors.RestError {
	time, err := date_utils.GetNowDate()
	if err != nil {
		fmt.Sprintf("Error formatting date Details: %s \n", err.Error())
		return errors.NewInternalServerError(fmt.Sprintf("Error formatting date Details: %s", err.Error()))
	}
	user.DateCreated = time

	er := users_db.UsersDBClient.Ping()
	if er != nil {
		log.Panic(er)
	}

	stmt, err := users_db.UsersDBClient.Prepare(queryInsertUser)

	if err != nil {
		fmt.Sprintf("Error when preparing query Details: %s \n", err.Error())

		return errors.NewInternalServerError(fmt.Sprintf("Error when preparing query Details: %s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		fmt.Sprintf("Error when preparing query Details: %s \n", err.Error())
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when preparing query Details: %s", err.Error()))

	}
	lastID, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Sprintf("Unable to pick the last insert id. Detais %s \n", err.Error())
		return errors.NewInternalServerError(

			fmt.Sprintf("Unable to pick the last insert id. Detais %s", err.Error()))
	}
	user.Id = lastID
	// result := userDB[user.Id]
	// if result != nil {
	// 	if result.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("User  %v already exist", result.Email))

	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("User id %d already exist", user.Id))
	// }

	// log.Println(user)

	// userDB[user.Id] = user

	return nil
}
func (user *User) Get() *errors.RestError {
	er := users_db.UsersDBClient.Ping()
	if er != nil {
		log.Panic(er)
	}
	//Works just queries as though multiple records reuturned
	// rows, err := users_db.UsersDBClient.Query(querySelectUser, user.Id)
	// defer rows.Close()
	// if err != nil {
	// 	return errors.NewInternalServerError(fmt.Sprintf("Error querying dataset Details: %s", err.Error()))
	// }
	// defer rows.Close()

	// counter := 0
	// var abc string
	// for rows.Next() {
	// 	if counter == 0 {

	// 		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
	// 			return errors.NewInternalServerError(fmt.Sprintf("Error scanning dataset Details: %s", err.Error()))
	// 		}
	// 	}
	// 	counter += 1

	// }

	// if counter == 0 {
	// 	return errors.NewNotFoundError(fmt.Sprintf("User id: %d not found details %d", user.Id, counter))
	// }
	// return nil

	stmt, err := users_db.UsersDBClient.Prepare(querySelectUser)

	if err != nil {
		fmt.Sprintf("Error when preparing query Details: %s \n", err.Error())

		return errors.NewInternalServerError(fmt.Sprintf("Error when preparing query Details: %s", err.Error()))
	}
	defer stmt.Close()
	rows := stmt.QueryRow(user.Id)

	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error retrieving dataset Details: %s", err))
	}
	return nil
}
