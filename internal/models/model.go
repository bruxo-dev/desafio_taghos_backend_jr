package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Books    []Book `json:"books"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Synopsis string `json:"synopsis"`
	UserID   int    `json:"user_id"`
	User     User   `json:"user"`
}
