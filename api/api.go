package api

import (
	"FRSC-Project/service"
	"net/http"
)

type BusApi interface {
	RegisterBus(w http.ResponseWriter, r *http.Request)
	GetBuses(w http.ResponseWriter, r *http.Request)
	CheckMembership(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service service.BusService
}

func NewBusApi(service service.BusService) BusApi { return &api{service: service} }
