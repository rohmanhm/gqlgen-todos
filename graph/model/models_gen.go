// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
	User   *User  `json:"user"`
}

type TodoInput struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Email string `json:"email"`
}
