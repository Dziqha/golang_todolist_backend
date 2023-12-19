package controllers

import (
	"Todo/helpers"
	"Todo/models"
	"net/http"
	"strconv"

	_ "github.com/go-playground/locales/id"
	"github.com/labstack/echo/v4"
)

func FeacthallTodo(c echo.Context) error {
	//mengekstrak token menjadi data
	data, err := helpers.ExtractDataFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	// memindah atau mengquery id ke database
	ownerFloat, ok := data["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to extract user ID"})
	}
	//mengubah id float menjadi integer
	owner := int(ownerFloat)

	result, err := models.FetchAlltodo(owner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTodo(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	deadline := c.FormValue("deadline")

	data, err := helpers.ExtractDataFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	ownerFloat, ok := data["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to extract user ID"})
	}
	//mengubah id float menjadi integer
	owner := int(ownerFloat)

	// Panggil fungsi Storetodo
	result, err := models.Storetodo(owner, title, description, deadline)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the response from the Storetodo function
	return c.JSON(http.StatusOK, result)
}

func UpdateTodo(c echo.Context) error {
	id := c.FormValue("id")
	owner := c.FormValue("owner")
	title := c.FormValue("title")
	description := c.FormValue("description")
	deadline := c.FormValue("deadline")
	complete := c.FormValue("complete")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	conv_owner, err := strconv.Atoi(owner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	conv_complete, err := strconv.ParseBool(complete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Perubahan variabel dari `result` menjadi `err`
	err = models.UpdateTodo(conv_id, conv_owner, title, description, deadline, conv_complete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Data updated successfully",
	})
}

func DeleteTodo(c echo.Context) error {
	id := c.FormValue("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data, err := helpers.ExtractDataFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	ownerFloat, ok := data["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to extract user ID"})
	}
	//mengubah id float menjadi integer
	owner := int(ownerFloat)

	err = models.DeleteTodo(conv_id, owner)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Data Deleted successfully",
	})
}
