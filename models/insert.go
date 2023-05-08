package models

import (
	"todo/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return

	}

	defer conn.Close()

	sql := `INSERT INTO todos (titile, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Titile, todo.Description, todo.Done).Scan((&id))
	return
}
