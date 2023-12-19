package routers

import (
	"Todo/controllers"
	"Todo/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)
	e.PUT("/auth/update", controllers.Update, middleware.IsAuthenticated)

	e.GET("/todo/view", controllers.FeacthallTodo, middleware.IsAuthenticated)
	e.POST("/todo/create", controllers.StoreTodo, middleware.IsAuthenticated)
	e.PUT("/todo/update", controllers.UpdateTodo, middleware.IsAuthenticated)
	e.DELETE("/todo/delete", controllers.DeleteTodo, middleware.IsAuthenticated)

	return e
}
