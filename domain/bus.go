package domain

import (
	"FRSC-Project/model"
	//"context"
	"errors"
	"log"
	//"net/http"
)

func (d *domain) CheckMembership(bus *model.Motor) (*model.Motor, error) {
	return d.db.FindOne(bus)
}

func (d *domain) RegisterBus(input model.Motor) (*model.AuthResponse, error) {
	_, err := d.db.GetBusByPlateNumber(input.PlateNumber)
	if err == nil {
		return &model.AuthResponse{
			Success:    false,
			Message:    "A bus with the plate number already exist",
			StatusCode: 400,
			AuthToken:  nil,
			Motor:      nil,
		}, nil
	}

	payload := &model.Motor{
		Brand:       input.Brand,
		Color:       input.Color,
		Seats:       input.Seats,
		PlateNumber: input.PlateNumber,
		Membership:  input.Membership,
		Model:       input.Model,
	}

	_, err = d.db.AddBus(payload)
	if err != nil {
		log.Printf("error while registering user: %v", err)
		return nil, errors.New("something went wrong, please try again")
	}

	token, err := payload.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return &model.AuthResponse{
		Success:    true,
		Message:    "Registration Successful",
		StatusCode: 201,
		AuthToken:  token,
		Motor:      payload,
	}, nil
}

func (d *domain) Buses() ([]*model.Motor, error) {
	return d.db.GetBuses()
}
