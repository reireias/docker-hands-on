package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Task data
type Task struct {
	ID   int
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	dataSource := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + "/hello"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}

	tasks := []Task{}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Name)
		if err != nil {
			fmt.Fprintf(w, "Error")
			return
		}
		tasks = append(tasks, task)
	}

	fmt.Fprintf(w, "%v", tasks)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}
