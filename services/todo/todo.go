package todo_service

import (
	"errors"
	"fmt"
	"time"

	"github.com/9d4/tinywebserver/database"
	"github.com/9d4/tinywebserver/helper"
)

var tableName = "todos"

type Todo struct {
	Id        int
	Content   string
	Done      interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}

type List []Todo

func New(content string) (int, error) {
	if helper.IsEmptyString(content) {
		return -1, errors.New("content is empty")
	}

	db := database.DB

	query := `INSERT INTO ` + tableName + ` (content) VALUES (?)`
	stmt, err := db.Prepare(query)

	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(content)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func GetList() List {
	out := []Todo{}

	query := `SELECT * FROM ` + tableName + ` ORDER BY done ASC, created_at DESC`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return out
	}

	rows, err := stmt.Query()
	if err != nil {
		return out
	}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Content, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return out
		}

		out = append(out, todo)
	}

	for i := range out {
		todo := &out[i]
		t, _ := time.LoadLocation("Asia/Jakarta")

		todo.CreatedAt = todo.CreatedAt.In(t)
		todo.UpdatedAt = todo.UpdatedAt.In(t)
	}

	return out
}

func Delete(id int) error {
	query := `DELETE FROM ` + tableName + ` WHERE id=(?)`
	stmt, _ := database.DB.Prepare(query)
	_, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func SetStatus(id int, status string) error {
	var statusInt int

	switch status {
	case "DONE":
		statusInt = 1
		break

	case "UNDONE":
		statusInt = 0
		break

	default:
		return nil
	}

	query := `UPDATE ` + tableName + ` SET done=(?) WHERE id=(?)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	stmt.Exec(statusInt, id)

	return nil
}
