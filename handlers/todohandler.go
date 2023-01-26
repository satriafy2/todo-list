package handlers

import (
	"fmt"
	"net/http"
	"todo-list/db"
	"todo-list/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	mTodo  model.Todo
	mTodos []model.Todo
	res    *gorm.DB
)

//	GetTodo godoc
//	@Summary Get all todo list
//	@Tags Todo
//	@Accept */*
//	@Produce json
//	@Success 200 {object} []model.Todo
//	@Router /todo [get]
func GetTodo(c echo.Context) error {
	showChild := c.QueryParam("show_child")
	todoID := c.Param("id")
	mTodos = []model.Todo{}

	if todoID != "" {
		res = db.Conn.Where("id=?", todoID).Find(&mTodos)
		showChild = "1"
	} else {
		res = db.Conn.Where("parent_id IS NULL").Find(&mTodos)
	}

	if res.Error != nil {
		fmt.Println("error query: ", res.Error)
		return c.NoContent(http.StatusInternalServerError)
	}
	dataResponse := populateParentTodoResponse(&mTodos, showChild)

	return c.JSON(http.StatusOK, dataResponse)
}

//	GetTodoSingle godoc
//	@Summary Get single todo
//	@Tags Todo
//	@Accept */*
//	@Produce json
//	@Success 200 {object} []model.Todo
//	@Router /todo/{id} [get]
func GetTodoSingle(c echo.Context) error {
	todoID := c.Param("id")
	res = db.Conn.Where("id=?", todoID).Find(&mTodos)

	if res.Error != nil {
		fmt.Println("error query: ", res.Error)
		return c.NoContent(http.StatusInternalServerError)
	}
	dataResponse := populateParentTodoResponse(&mTodos, "1")

	return c.JSON(http.StatusOK, dataResponse)
}

//	CreateTodo godoc
//	@Summary Create to-do list
//	@Tags Todo
//	@Accept json
//	@Produce json
//	@Param todo body CreateTodoRequest true "Create Todo Body, include parent_id field to create sub todo list"
//	@Success 201 {object} model.Todo
//	@Router /todo [post]
func CreateTodo(c echo.Context) error {
	t := new(CreateTodoRequest)
	if err := c.Bind(t); err != nil {
		return err
	}

	mTodo = model.Todo{}
	err := t.populate(&mTodo)
	if err != nil {
		return err
	}

	res = db.Conn.Create(&mTodo)
	if res.Error != nil {
		fmt.Println("error query: ", res.Error)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, mTodo)
}

//	UpdateTodo godoc
//	@Summary Update to-do
//	@Tags Todo
//	@Accept json
//	@Produce json
//	@Param todo body UpdateTodoRequest true "Update todo request body, id field is required"
//	@Success 201 {object} model.Todo
//	@Router /todo/{id} [patch]
func UpdateTodo(c echo.Context) error {
	todoID := c.Param("id")
	t := new(UpdateTodoRequest)
	if err := c.Bind(t); err != nil {
		return err
	}

	res = db.Conn.First(&mTodo, todoID)
	if res.RowsAffected == 0 {
		return c.NoContent(http.StatusBadRequest)
	}
	t.populate(&mTodo)

	res = db.Conn.Save(&mTodo)
	if res.Error != nil {
		fmt.Println("error query: ", res.Error)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, mTodo)
}

//	DeleteTodo godoc
//	@Summary Delete to-do
//	@Tags Todo
//	@Accept */*
//	@Produce text/plain
//	@Success 202
//	@Router /todo/{id} [delete]
func DeleteTodo(c echo.Context) error {
	todoID := c.Param("id")
	res = db.Conn.Where("id=? OR parent_id=?", todoID, todoID).Find(&mTodos)
	if res.RowsAffected == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	res = db.Conn.Delete(&mTodos)
	if res.Error != nil {
		return res.Error
	}

	return c.NoContent(http.StatusAccepted)
}
