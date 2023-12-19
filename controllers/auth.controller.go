package controllers

import (
	"Todo/helpers"
	"Todo/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Panggil fungsi Checklogin dari models
	user, err := models.Checklogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// Jika user nil, berarti autentikasi gagal
	if user == nil {
		return echo.ErrUnauthorized
	}

	// Jika autentikasi berhasil, buat token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("bisa"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.Registermasuk(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Update(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	data, err := helpers.ExtractDataFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	pemilikFloat, ok := data["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to extract user ID"})
	}

	pemilik := int(pemilikFloat)

	result, err := models.Updatelogin(username, password, pemilik)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
