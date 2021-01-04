package models

import (
    "mechanics-backend/persistance"
)

type Appointment struct {
    Id int64
    VehicleId int64
    Date int64
}

func (a *Appointment) GetId() int64 {
    return a.Id
}

// TODO save the retrieved vehicle into a private field
func (a *Appointment) Vehicle (repo *persistance.Repository) *Vehicle  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, a.VehicleId)
    return vehicle
}

func (a *Appointment) Client (repo *persistance.Repository) *Client  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, a.VehicleId)
    client := vehicle.Client(repo)
    return client
}