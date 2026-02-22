package db

import "time"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Pages  int    `json:"pages"`
}

type User struct {
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
