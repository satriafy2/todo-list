package handlers

import (
	"time"
	"todo-list/db"
	"todo-list/model"
)

type CreateTodoRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	ParentID    *int   `json:"parent_id" form:"parent_id"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	ParentID    *int    `json:"parent_id" form:"parent_id"`
}

func (r *CreateTodoRequest) populate(t *model.Todo) error {
	t.Title = r.Title
	t.Description = r.Description
	t.CreatedAt = time.Now()

	if r.ParentID != nil {
		result := db.Conn.Where("id=? AND parent_id IS NULL", r.ParentID).First(&mTodos)
		if result.Error != nil {
			return result.Error
		}
		t.ParentID = *r.ParentID
	}

	return nil
}

func (r *UpdateTodoRequest) populate(t *model.Todo) error {
	if r.Title != nil {
		t.Title = *r.Title
	}
	if r.Description != nil {
		t.Description = *r.Description
	}
	t.UpdatedAt = time.Now()

	if r.ParentID != nil {
		result := db.Conn.Where("id=? AND parent_id IS NULL", r.ParentID).First(&mTodos)
		if result.Error != nil {
			return result.Error
		}
		t.ParentID = *r.ParentID
	}

	return nil
}
