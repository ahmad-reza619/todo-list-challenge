package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Activity struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt NullTime  `json:"deleted_at"`
}

func FindAllActivity(db *sql.DB) []Activity {
	activityRec, err := db.Query("select id, title, email, created_at, updated_at, deleted_at from activity")
	if err != nil {
		panic(err.Error())
	}
	activities := []Activity{}
	for activityRec.Next() {
		activity := Activity{}
		err = activityRec.Scan(&activity.Id, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)

		if err != nil {
			panic(err.Error())
		}

		activities = append(activities, activity)
	}
	return activities
}

func FindByActivityId(db *sql.DB, id int64) Activity {
	activityRec, err := db.Query("select id, title, email, created_at, updated_at, deleted_at from activity where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	activity := Activity{}
	for activityRec.Next() {
		err = activityRec.Scan(&activity.Id, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)

		if err != nil {
			panic(err.Error())
		}
	}

	return activity
}

func AddActivity(db *sql.DB, title string, email string) int64 {
	insOps, err := db.Prepare("INSERT INTO activity(title, email) VALUES(?, ?)")

	if err != nil {
		panic(err.Error())
	}
	res, err := insOps.Exec(title, email)
	if err != nil {
		panic(err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return lastId
}

func DeleteActivityById(db *sql.DB, id int64) int64 {
	delOps, err := db.Prepare("DELETE from activity where id = ?")

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
func UpdateActivityById(db *sql.DB, id int64, title string) {
	updateOps, err := db.Prepare("UPDATE activity SET title = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateOps.Exec(title, id)
}
