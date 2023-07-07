package models

type Todo struct {
	Id          int
	UserId      int `gorm:"user_id" json:"-"`
	Title       string
	Description string
	IsComplete  bool
}
type CreateTodoInput struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	IsComplete  bool   `json:"is_complete,omitempty" default:"false"`
}
