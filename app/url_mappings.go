package app

import (
	"users_api/controllers/ping"
	"users_api/controllers/users"
)

func map_Urls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.PUT("/users/:user_id", users.SearchUser)
}
