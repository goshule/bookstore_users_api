package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	UsersDBClient *sql.DB
)

func init() {

	Mysql_username := "mysql_username"
	Mysql_password := "mysql_password"
	Mysql_host := "mysql_host"
	Mysql_schema := "mysql_schema"

	// errrr := godotenv.Load("app/creds.env")
	// if errrr != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Mysql_username = os.Getenv("mysql_username")
	// Mysql_password = os.Getenv("mysql_password")
	// Mysql_host = os.Getenv("mysql_host")
	// Mysql_schema = os.Getenv("mysql_schema")

	Mysql_username = "root"
	Mysql_password = "root123"
	Mysql_host = "127.0.0.1"
	Mysql_schema = "bookstore_users"

	//log.Println("-------------------", Mysql_username, Mysql_host, Mysql_schema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", Mysql_username, Mysql_password, Mysql_host, Mysql_schema)

	var err error
	UsersDBClient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	er := UsersDBClient.Ping()

	if er != nil {
		log.Panic(err)
	}

	log.Println("Database successfully connected")

}
