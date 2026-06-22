package main

import (
	"fmt"
	"net/http"
)

func main() {
	// database
	db, err := openDB()
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = CreateListItemsTable(db)
	if err != nil {
		panic(err)
	}

	// app
	app := &App{
		DB: db,
	}

	// server
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/list", app.ListItemsHandler)

	fmt.Println("Server running on port:8888")

	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		panic(err)
	}
}
