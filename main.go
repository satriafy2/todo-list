package main

import (
	"todo-list/db"
	"todo-list/router"

	_ "todo-list/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title To-Do List API
// @version 1.0
// @description This is a to-do list mini project.
// @termsOfService http://swagger.io/terms/
// @host localhost:1323
// @BasePath /
func main() {
	db.Init()
	database, _ := db.Conn.DB()
	db.AutoMigrate(db.Conn)
	defer database.Close()

	e := router.Init()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
