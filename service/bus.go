package service

import (
	"FRSC-Project/model"
	"errors"
)

func (s *service) Validate(motor *model.Motor) error {
	if motor == nil {
		err := errors.New("the post is empty")
		return err
	}

	if motor.Brand == "" {
		err := errors.New("the bus brand name can not be empty")
		return err
	}

	if motor.PlateNumber == "" {
		err := errors.New("the bus plate number can not be empty")
		return err
	}

	if motor.Membership == "" {
		err := errors.New("please provide your bus membership")
		return err
	}

	if motor.Seats == "" {
		err := errors.New("please provide seat capacity")
		return err
	}

	if motor.Color == "" {
		err := errors.New("please provide the color of your bus")
		return err
	}
	return nil
}

func (s *service) CheckMembership(bus *model.Motor) (*model.Motor, error) {
	return s.domain.CheckMembership(bus)
}

func (s *service) RegisterBus(input model.Motor) (*model.AuthResponse, error) {
	return s.domain.RegisterBus(input)
}

func (s *service) Buses() ([]*model.Motor, error) {
	return s.domain.Buses()
}
