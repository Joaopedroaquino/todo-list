package models

type Todo struct {
	ID          int64  `json: "id"`
	Titile      string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}
