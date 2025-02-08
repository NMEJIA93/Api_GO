package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"
	http.HandleFunc("/user", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /users")
	io.WriteString(w, "This is My user endpoint!\n")
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /courses")
	io.WriteString(w, "This is My course endpoint!\n")
}
