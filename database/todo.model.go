package database

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type Todo struct {
	Id              int64     `json:"id"`
	ActivityGroupId int64     `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       NullTime  `json:"deleted_at"`
}

func FindAllTodo(db *sql.DB) []Todo {
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todos")
	if err != nil {
		panic(err.Error())
	}
	todos := []Todo{}
	for todoRec.Next() {
		todo := Todo{}
		err = todoRec.Scan(&todo.Id, &todo.ActivityGroupId, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

		if err != nil {
			panic(err.Error())
		}

		todos = append(todos, todo)
	}
	return todos
}

func FindTodoById(db *sql.DB, id int64) (Todo, error) {
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todos where id=?", id)
	defer todoRec.Close()
	if err != nil {
		return Todo{}, err
	}
	todo := Todo{}
	if todoRec.Next() {
		err = todoRec.Scan(&todo.Id, &todo.ActivityGroupId, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

		if err != nil {
			return Todo{}, err
		}
	} else {
		return Todo{}, errors.New("No Records Found")
	}

	return todo, nil
}

func FindTodoByActivityId(db *sql.DB, id int64) []Todo {
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todos where activity_group_id=?", id)
	defer todoRec.Close()
	if err != nil {
		panic(err.Error())
	}
	todos := []Todo{}
	for todoRec.Next() {
		todo := Todo{}
		err := todoRec.Scan(&todo.Id, &todo.ActivityGroupId, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

		if err != nil {
			panic(err.Error())
		}

		todos = append(todos, todo)
	}

	return todos
}

func AddTodo(db *sql.DB, title string, activity_group_id int64) int64 {
	insOps, err := db.Prepare("INSERT INTO todos(title, activity_group_id, is_active, priority) VALUES(?, ?, ?, ?)")

	defer insOps.Close()
	if err != nil {
		panic(err.Error())
	}
	res, err := insOps.Exec(title, activity_group_id, "1", "very-high")
	if err != nil {
		panic(err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return lastId
}

func DeleteTodoById(db *sql.DB, id int64) error {
	delOps, err := db.Prepare("DELETE from todos where id = ?")
	defer delOps.Close()

	if err != nil {
		return err
	}
	res, err := delOps.Exec(id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows > 0 {
		return nil
	}
	return errors.New("No Records Found")
}
func UpdateTodoById(db *sql.DB, id int64, title string, isActive *bool) {
	q := `UPDATE todos SET `
	qParts := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)

	if title != "" {
		qParts = append(qParts, `title = ?`)
		args = append(args, title)
	}
	if isActive != nil {
		qParts = append(qParts, `is_active = ?`)
		args = append(args, isActive)
	}
	q += strings.Join(qParts, `,`) + ` WHERE id = ?`
	args = append(args, id)
	updateOps, err := db.Prepare(q)
	defer updateOps.Close()
	if err != nil {
		panic(err.Error())
	}
	updateOps.Exec(args...)
}
