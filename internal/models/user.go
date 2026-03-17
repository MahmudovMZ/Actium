package models

type User struct {
	ID        int    `json:"-"`
	UserName  string `json:"user_name"`
	Password  string `json:"user_pass"`
	CreatedAt string `json:"created_at"`
}
