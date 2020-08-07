package api

import (
	"FRSC-Project/model"
	"encoding/json"
	"log"
	"net/http"
)

func (a *api) RegisterBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buses model.Motor
	err := json.NewDecoder(r.Body).Decode(&buses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
