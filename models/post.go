package models

type Post struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Preview string `json:"preview"`
	Content string `json:"content"`
}
