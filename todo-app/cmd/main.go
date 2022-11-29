package main

import (
	"github.com/geges1101/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured running http server: %s", err.Error())
	}
}
