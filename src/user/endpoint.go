package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}

}

func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
