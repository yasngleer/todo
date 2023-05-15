package types

import "time"

type TodoList struct {
	ID                  int
	Name                string
	User                User
	UserID              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
	PercentOfCompletion int
	Steps               []TodoStep `gorm:"foreignKey:TodoListID"`
}
type TodoStep struct {
	ID         int
	Context    string
	TodoListID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Completed  bool
}
type TodoListResponse struct {
	ID                  int
	Name                string
	Author              UserResponse
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
	Steps               []TodoStep
	PercentOfCompletion int
}

func (t *TodoList) ToResponse() *TodoListResponse {
	return &TodoListResponse{
		ID:                  t.ID,
		Name:                t.Name,
		Author:              *t.User.ToResponse(),
		CreatedAt:           t.CreatedAt,
		UpdatedAt:           t.UpdatedAt,
		DeletedAt:           t.DeletedAt,
		Steps:               t.Steps,
		PercentOfCompletion: t.PercentOfCompletion,
	}
}
