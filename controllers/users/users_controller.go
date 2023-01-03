package users

import (
	"net/http"
	"users_api/domain/users"
	"users_api/services"
	"users_api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO:Handle error
	// 	return
	// }
	// e := json.Unmarshal(bytes, &user)
	// if e != nil {
	// 	//TODO:Handle error
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid JSON body")

		c.JSON(restError.Status, restError)
		// fmt.Println(restError)
		return
	}

	saveRes, er := services.CreateUser(&user)
	if er != nil {
		c.JSON(er.Status, er)
		return
	}
	c.JSON(http.StatusCreated, saveRes)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
