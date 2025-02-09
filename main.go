package main

import (
	"log"
	"net/http"
	"time"

	"github.com/NMEJIA93/Api_GO/src/user"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	userService := user.NewService()

	userEnd := user.MakeEndpoints(userService)

	router.HandleFunc("/user", userEnd.Get).Methods("GET")
	router.HandleFunc("/user", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/user", userEnd.Create).Methods("POST")
	router.HandleFunc("/user", userEnd.Update).Methods("PUT")
	router.HandleFunc("/user", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
