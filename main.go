package main

import (
	"./api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()

	if err != nil {
		panic("Error opening log file")
	}
	log.SetOutput(f)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/items", api.GetItems).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))

	http.Handle("/", router)

	fmt.Println("Application running on port 8080")
	http.ListenAndServe(":8080", nil)
}
