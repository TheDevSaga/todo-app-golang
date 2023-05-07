package models

type Todo struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
}
type CreateTodoInput struct {
	Title       string
	Description string
	IsComplete  bool
}
