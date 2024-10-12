package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx context.Context
}
type Todo struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	COMPLETED int    `json:"completed"`
}

var DB *sql.DB

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		completed INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	a.ctx = ctx
}

func (a *App) CreateEntry(name string) bool {
	if DB == nil {
		log.Fatal("DB illadoi")
		return false
	}
	_, err := DB.Exec("INSERT INTO todos (name, completed) VALUES (?, ?)", name, 0)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
func (a *App) UpdateEntry(id int, name string, completed int) bool {
	if DB == nil {
		log.Fatal("DB illadoi")
		return false
	}
	_, err := DB.Exec("UPDATE todos SET completed = ?, name = ? WHERE id = ?", completed, name, id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
func (a *App) DeleteEntry(id int) bool {
	if DB == nil {
		log.Fatal("DB illadoi")
		return false
	}
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (a *App) ReadData() []Todo {
	var results []Todo
	if DB == nil {
		log.Fatal("DB illadoi")
		return results
	}
	readSQL := `SELECT id, name, completed FROM todos`
	rows, err := DB.Query(readSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.NAME, &todo.COMPLETED)
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", todo.ID, todo.NAME, todo.COMPLETED)
		results = append(results, todo)
	}
	log.Println("Slice:", results)
	return results
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	defer DB.Close()
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
