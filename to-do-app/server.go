package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danierj/training/to-do-app/todo"
)

func initServer(todos []todo.Todo) {

	port := ":8080"

	fmt.Println("\n-------------------- Starting server at port " + port + " --------------------")

	http.HandleFunc("/todos", getTodosHandler(todos))

	log.Fatal(http.ListenAndServe(port, nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func getTodosHandler(todos []todo.Todo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")

		json, err := json.MarshalIndent(todos, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		w.Write(json)

	}
}
