package main

import (
	"log"
	"net/http"
	"time"

	"github.com/NMEJIA93/Api_GO/pkg/bootstrap"
	"github.com/NMEJIA93/Api_GO/src/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	_ = godotenv.Load()

	l := bootstrap.InitLogger()

	db, err := bootstrap.BDConnection()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(l, db)
	userService := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userService)

	router.HandleFunc("/user/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/user", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/user", userEnd.Create).Methods("POST")
	router.HandleFunc("/user", userEnd.Update).Methods("PUT")
	router.HandleFunc("/user/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/user/{id}", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err1 := srv.ListenAndServe()
	if err1 != nil {
		log.Fatal(err1)
	}
}
