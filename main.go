package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NMEJIA93/Api_GO/src/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := mux.NewRouter()

	_ = godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	fmt.Println(dsn)
	db, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err1 != nil {
		log.Fatalf("failed to initialize database, got error %v", err1)
	}

	er := db.Debug().AutoMigrate(&user.User{})
	if er != nil {
		log.Fatal(er)
	}

	er = db.AutoMigrate(&user.User{})

	if er != nil {
		log.Fatal(er)
	}

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
