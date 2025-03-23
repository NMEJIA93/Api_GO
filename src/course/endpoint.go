package course

import (
	"encoding/json"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: "Invalid request", Data: nil})
			return
		}
		if req.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: "name is required", Data: nil})
			return
		}

		if req.StartDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: "start date is required", Data: nil})
			return
		}
		if req.EndDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: "end date is required", Data: nil})
			return
		}
		dto := CreateCourseDTO{
			Name:      req.Name,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		}
		course, err := s.Create(dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: err.Error()})
			return

		}
		json.NewEncoder(w).Encode(&Response{Status: 200, Data: course})

	}
}
