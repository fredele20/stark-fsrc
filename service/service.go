package service

import (
	"FRSC-Project/domain"
	"FRSC-Project/model"
)

type BusService interface {
	RegisterBus(input model.Motor) (*model.AuthResponse, error)
	Buses() ([]*model.Motor, error)
	Validate(motor *model.Motor) error
	CheckMembership(bus *model.Motor) (*model.Motor, error)
}

type service struct {
	domain domain.BusDomain
}

func NewBusService(domain domain.BusDomain) BusService { return &service{domain} }
