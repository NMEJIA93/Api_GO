package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	ErrorResp struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"Invalid request"})
			return
		}

		if req.FirstName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"first name is required"})
		}
		if req.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"first name is required"})
		}

		dto := CreateUserDTO{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		user, serviceErr := s.Create(dto)
		if serviceErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{serviceErr.Error()})
			return
		}
		responseDto := ResponseUserDto{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		}

		fmt.Println("Create user: ", responseDto.ID)
		json.NewEncoder(w).Encode(responseDto)
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		err := s.Delete(id)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{"User not found"})
			return
		}

		fmt.Println("Delete user")
		json.NewEncoder(w).Encode(map[string]string{"data": "Success"})
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		user, err := s.Get(id)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{err.Error()})
			return
		}
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(user)
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All user")
		users, err := s.GetAll()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{err.Error()})
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update user")

		var req UpdateReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"Invalid request"})
			return
		}

		if req.FirstName != nil && *req.FirstName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"first name is required"})
			return
		}

		if req.FirstName != nil && *req.LastName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"last name is required"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		err = s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{"User not found"})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"data": "success"})
	}
}
