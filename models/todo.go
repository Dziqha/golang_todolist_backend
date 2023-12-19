package models

import (
	"Todo/db"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Todolist struct {
	Id          int    `json:"id"`
	Owner       int    `json:"owner" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Deadline    string `json:"deadline" validate:"required"`
	Complete    bool   `json:"complete" `
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FetchAlltodo(owner int) (AuthResponse, error) {
	var obj Todolist
	var arrobj []Todolist
	var res AuthResponse

	// Get a database connection
	con := db.CreateCon()
	defer con.Close()

	// Perform the SQL query with a WHERE clause to filter by 'pemilik'
	sqlStatement := "SELECT * FROM todos WHERE owner = ?"
	rows, err := con.Query(sqlStatement, owner)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return res, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		// Scan each row into the 'obj' variable
		err := rows.Scan(&obj.Id, &obj.Owner, &obj.Title, &obj.Description, &obj.Deadline, &obj.Complete)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return res, err
		}
		// Append the scanned object to the result array
		arrobj = append(arrobj, obj)
	}

	// Populate the response structure
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func Storetodo(owner int, title string, description string, deadline string) (AuthResponse, error) {
	var res AuthResponse

	v := validator.New()

	tod := Todolist{
		Owner:       owner,
		Title:       title,
		Description: description,
		Deadline:    deadline,
	}

	err := v.Struct(tod)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	// Kolom ID diabaikan karena diatur sebagai AUTO_INCREMENT
	sqlStatement := "INSERT INTO todos (owner, title, description, deadline) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(owner, title, description, deadline)
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

func UpdateTodo(id int, owner int, title string, description string, deadline string, complete bool) error {
	var res AuthResponse
	con := db.CreateCon()

	sqlstatement := "UPDATE todos SET owner = ?, title = ?, description = ?, deadline = ?, complete = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(owner, title, description, deadline, complete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return nil
}
func DeleteTodo(id int, owner int) error {
	var res AuthResponse
	con := db.CreateCon()
	defer con.Close()

	sqlstatement := "DELETE FROM todos WHERE id =? AND owner = ?"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(id, owner)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected < 1 {
		return fmt.Errorf("no rows affected, todo not found")
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return nil
}
