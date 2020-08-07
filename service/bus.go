package service

import "FRSC-Project/model"

func (s *service) RegisterBus(input model.Motor) (*model.AuthResponse, error) {
	return s.domain.RegisterBus(input)
}

func (s *service) Buses() ([]*model.Motor, error) {
	return s.domain.Buses()
}
