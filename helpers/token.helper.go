package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ExtractDataFromToken(c echo.Context) (map[string]interface{}, error) {
	tokenString := c.Request().Header.Get("Authorization")

	// Check if token is missing
	if tokenString == "" {
		return nil, errors.New("token is missing")
	}

	// Token is usually included in the header in the format "Bearer <token>"
	splitToken := strings.Split(tokenString, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("invalid token format")
	}

	tokenString = splitToken[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You need to return the secret key used to sign the token
		return []byte("bisa"), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	// Check token validity
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to extract claims")
	}

	// Extract data from claims
	data := make(map[string]interface{})
	for key, value := range claims {
		data[key] = value
	}

	return data, nil
}
