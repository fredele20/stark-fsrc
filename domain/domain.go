package domain

import (
	"FRSC-Project/database"
	"FRSC-Project/model"
)

type BusDomain interface {
	RegisterBus(input model.Motor) (*model.AuthResponse, error)
	Buses() ([]*model.Motor, error)
}

type domain struct {
	db database.Bus
}

func NewBusDomain(dbs database.Bus) BusDomain { return &domain{db: dbs} }
