package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/todo-list?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

type Activity struct {
	Id        int64
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt mysql.NullTime
}

func findActivityById(db *sql.DB, id int64) Activity {
	activityRec, err := db.Query("select * from activity where id=?", id)
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

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(struct {
		Text string `json:"text"`
	}{
		"Hello World",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(data)
}

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()
	if r.Method == "POST" {
		title, email := r.FormValue("title"), r.FormValue("email")

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
		log.Println("ADDED ACTIVITY: title{" + title + "} email{" + email + "} with ID " + strconv.FormatInt(lastId, 10))
		activity := findActivityById(db, lastId)

		type DataActivity struct {
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Id        int64     `json:"id"`
			Title     string    `json:"title"`
			Email     string    `json:"email"`
		}

		type ResponseActivity struct {
			Status  string       `json:"status"`
			Message string       `json:"message"`
			Data    DataActivity `json:"data"`
		}

		data := DataActivity{
			activity.CreatedAt,
			activity.UpdatedAt,
			activity.Id,
			activity.Title,
			activity.Email,
		}

		toJson := ResponseActivity{
			"Success",
			"Success",
			data,
		}

		response, err := json.Marshal(toJson)

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	}
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started at port 8081")

	http.HandleFunc("/", Index)
	http.HandleFunc("/activity-groups", CreateActivity)
	http.ListenAndServe(":8081", nil)
}
