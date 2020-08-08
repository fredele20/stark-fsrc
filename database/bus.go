package database

import (
	"FRSC-Project/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	//"log"
)

type Bus interface {
	AddBus(bus *model.Motor) (*model.Motor, error)
	GetBusByField(field, value string) (*model.Motor, error)
	GetBusByPlateNumber(plateNumber string) (*model.Motor, error)
	GetBuses() ([]*model.Motor, error)
	FindOne(bus *model.Motor) (*model.Motor, error)
	//GetBus()
}

type db struct {
}

func NewBusDB() Bus { return &db{} }

func (d *db) FindOne(bus *model.Motor) (*model.Motor, error) {
	var buses *model.Motor
	err := Collection.FindOne(context.TODO(), bson.D{{"model", bus.Model}}).Decode(&buses)
	if err == nil {
		log.Fatalf("Could not get the post %v", err)
	}
	return buses, nil
}

func (d *db) GetBusByField(field, value string) (*model.Motor, error) {
	var motor model.Motor
	err := Collection.FindOne(context.TODO(), bson.M{field: value}).Decode(&motor)
	return &motor, err
}

func (d *db) GetBusByPlateNumber(plateNumber string) (*model.Motor, error) {
	return d.GetBusByField("platenumber", plateNumber)
}

func (d *db) GetBuses() ([]*model.Motor, error) {
	var buses []*model.Motor

	cursor, err := Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var bus *model.Motor
		cursor.Decode(&bus)
		buses = append(buses, bus)
	}
	return buses, nil
}

func (d *db) AddBus(bus *model.Motor) (*model.Motor, error) {
	_, err := Collection.InsertOne(context.TODO(), bus)
	return bus, err
}
