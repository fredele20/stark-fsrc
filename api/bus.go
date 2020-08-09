package api

import (
	"FRSC-Project/error"
	"FRSC-Project/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}

func (a *api) RegisterBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buses model.Motor
	err := json.NewDecoder(r.Body).Decode(&buses)
	setupResponse(&w, r)
	if (*r).Method == "POST" {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = a.service.Validate(&buses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error.ServiceError{Message: err.Error()})
			return
		}

		bus, err := a.service.RegisterBus(buses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		//w.WriteHeader(model.StatusCreated)
		json.NewEncoder(w).Encode(bus)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("error")
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Method not allowed"})
	}
}

func (a *api) GetBuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	buses, err := a.service.Buses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buses)
}

func (a *api) CheckMembership(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buses model.Motor

	bus, err := a.service.CheckMembership(&buses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bus)
}
