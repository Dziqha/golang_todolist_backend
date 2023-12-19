package models

import (
	"Todo/db"
	"Todo/helpers"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
}

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Update struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Checklogin(username, password string) (*Users, error) {
	var user Users
	var hashedPassword string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&user.Id, &user.Username, &hashedPassword,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return nil, err
	}

	if err != nil {
		fmt.Println("Query error")
		return nil, err
	}

	match, err := helpers.CheckPassword(password, hashedPassword)
	if !match {
		fmt.Println("Password and hash don't match")
		return nil, err
	}

	return &user, nil
}

func Updatelogin(username string, password string, id int) (AuthResponse, error) {
	var res AuthResponse

	v := validator.New()

	up := Update{
		Username: username,
		Password: password,
	}
	err := v.Struct(up)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()

	sqlstatement := "UPDATE users SET username = ?, password = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}

	hash, _ := helpers.HashPassword(password)

	result, err := stmt.Exec(username, hash, id)

	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"lastInsertId": lastInsertId,
	}

	return res, nil
}

func Registermasuk(username string, password string) (AuthResponse, error) {
	var res AuthResponse

	v := validator.New()

	reg := Register{
		Username: username,
		Password: password,
	}
	err := v.Struct(reg)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()

	sqlstatement := "INSERT INTO users (username, password) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}

	hash, _ := helpers.HashPassword(password)

	result, err := stmt.Exec(username, hash)

	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"lastInsertId": lastInsertId,
	}

	return res, nil
}
