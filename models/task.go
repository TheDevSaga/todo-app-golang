package models

type Todo struct {
	Id          int
	UserId      int `gorm:"user_id" `
	Title       string
	Description string
	IsComplete  bool
}
type CreateTodoInput struct {
	Title       string
	Description string
	IsComplete  bool
}
