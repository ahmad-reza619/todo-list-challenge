package database

import (
	"database/sql"
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
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todo")
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

func FindTodoById(db *sql.DB, id int64) Todo {
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todo where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	todo := Todo{}
	for todoRec.Next() {
		err = todoRec.Scan(&todo.Id, &todo.ActivityGroupId, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

		if err != nil {
			panic(err.Error())
		}
	}

	return todo
}

func FindTodoByActivityId(db *sql.DB, id int64) []Todo {
	todoRec, err := db.Query("select id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at from todo where activity_group_id=?", id)
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
	insOps, err := db.Prepare("INSERT INTO todo(title, activity_group_id, is_active, priority) VALUES(?, ?, ?, ?)")

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

func DeleteTodoById(db *sql.DB, id int64) int64 {
	delOps, err := db.Prepare("DELETE from todo where id = ?")

	if err != nil {
		panic(err.Error())
	}
	res, err := delOps.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	return rows
}
func UpdateTodoById(db *sql.DB, id int64, title string) {
	updateOps, err := db.Prepare("UPDATE todo SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateOps.Exec(title, id)
}
