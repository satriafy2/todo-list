package handlers

import (
	"time"
	"todo-list/db"
	"todo-list/model"
)

type TodoResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

type ParentTodoResponse struct {
	TodoResponse
	SubTodo *[]TodoResponse `json:"sub_todos,omitempty"`
}

func populateTodoResponse(t *[]model.Todo) *[]TodoResponse {
	todosResponse := []TodoResponse{}
	for _, todo := range *t {
		todosResponse = append(todosResponse, TodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Created:     todo.CreatedAt,
			Updated:     todo.UpdatedAt,
		})
	}

	return &todosResponse
}

func populateParentTodoResponse(t *[]model.Todo, s string) *[]ParentTodoResponse {
	todosResponse := []ParentTodoResponse{}
	for _, todo := range *t {
		parentResponse := ParentTodoResponse{
			TodoResponse: TodoResponse{
				ID:          todo.ID,
				Title:       todo.Title,
				Description: todo.Description,
				Created:     todo.CreatedAt,
				Updated:     todo.UpdatedAt,
			},
		}

		if s == "1" {
			subTodoResponse := []TodoResponse{}
			res := db.Conn.Where("parent_id=?", todo.ID).Find(&mTodos)
			if res.RowsAffected > 0 {
				subTodoResponse = *populateTodoResponse(&mTodos)
			}
			parentResponse.SubTodo = &subTodoResponse
		}

		todosResponse = append(todosResponse, parentResponse)
	}

	return &todosResponse
}
