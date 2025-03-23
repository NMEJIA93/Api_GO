package course

import (
	"encoding/json"
	meta2 "github.com/NMEJIA93/Api_GO/pkg/meta"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create  Controller
		GetById Controller
		GetAll  Controller
		Detele  Controller
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
		Meta   *meta2.Meta `json:"meta,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetById: makeGetByIdEndpoint(s),
		GetAll:  makeGetAllEndpoint(s),
		Detele:  makeDeleteEndpoint(s),
	}
}
func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		err := s.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				Status: 400,
				Error:  err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&Response{
			Status: 200,
			Data:   "Success",
		})
	}
}

func makeGetByIdEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		course, err := s.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&Response{Status: 200, Data: course})
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		v := r.URL.Query()
		filters := Filters{
			Name: v.Get("name"),
		}
		limit, _ := strconv.Atoi(v.Get("limit"))
		page, _ := strconv.Atoi(v.Get("page"))

		count, err := s.Count(filters)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 500, Error: err.Error()})
			return
		}
		meta2, err := meta2.New(page, limit, count)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 500, Error: err.Error()})
			return
		}

		courses, err := s.GetAll(filters, meta2.Offset(), meta2.Limit())

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{Status: 400, Error: err.Error()})
			return
		}
		json.NewEncoder(w).Encode(&Response{
			Status: 200,
			Data:   courses,
			Meta:   meta2})
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
