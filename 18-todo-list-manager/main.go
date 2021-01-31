package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func init() {
	database, err := sql.Open("sqlite3", "database")

	if err != nil {
		log.Fatal(err)
	} else {
		defer database.Close()
	}
}

func main() {
	getTaskSQL := "select * from task where finish_date is null"

	rows, err := database.Query(getTaskSQL)

	if err != nil {
		log.Fatal(err)
	} else {
		defer rows.Close()
	}

	for rows.Next() {
		var Id int
		var Title string

		err := rows.Scan(&Id, &Title)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(Id, Title)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	const PORT = ":8080"

	// http.HandleFunc("/complete/", CompleteTaskFunc)
	// http.HandleFunc("/delete/", DeleteTaskFunc)
	// http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	// http.HandleFunc("/trash/", TrashTaskFunc)
	// http.HandleFunc("/edit/", EditTaskFunc)
	// http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	// http.HandleFunc("/restore/", RestoreTaskFunc)
	// http.HandleFunc("/add/", AddTaskFunc)
	// http.HandleFunc("/update/", UpdateTaskFunc)
	// http.HandleFunc("/search/", SearchTaskFunc)
	// http.HandleFunc("/login", GetLogin)
	// http.HandleFunc("/register", PostRegister)
	// http.HandleFunc("/admin", HandleAdmin)
	// http.HandleFunc("/add_user", PostAddUser)
	// http.HandleFunc("/change", PostChange)
	// http.HandleFunc("/logout", HandleLogout)
	http.HandleFunc("/", ShowAllTasksFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))

	log.Printf("Server is running on %s", PORT)

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string

	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}

	w.Write([]byte(message))
}
