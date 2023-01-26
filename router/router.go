package router

import (
	"net/http"
	"todo-list/handlers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome.")
	})

	e.GET("/todo", handlers.GetTodo)
	e.POST("/todo", handlers.CreateTodo)

	e.GET("/todo/:id", handlers.GetTodoSingle)
	e.PATCH("/todo/:id", handlers.UpdateTodo)
	e.DELETE("/todo/:id", handlers.DeleteTodo)

	return e
}
