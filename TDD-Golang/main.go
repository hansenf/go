package main

import (
	"log"
	"net/http"
)

func main() {
	repo := &employees.Repository{}
	_, err := repo.InitDB()
	if err != nil {
		log.Panic(err)
	}

	employees.NewRouter()
	http.ListenAndServe(":8080", nil)
}
