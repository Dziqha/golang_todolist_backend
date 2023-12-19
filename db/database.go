package db

import (
	"database/sql"
	"fmt"

	"Todo/config"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCon() *sql.DB {
	conf := config.GetConfig()

	dbDriver := "mysql"
	dbUser := conf.DB_USERNAME
	dbPass := conf.DB_PASSWORD
	dbName := conf.DB_NAME

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println("Error creating DB connection:", err)
		panic(err.Error())
	}
	return db
}
