package models

import (
    "mechanics-backend/persistance"
)

type Appointment struct {
    Id int64
    VehicleId int64
    Date int64
}

// TODO save the retrieved vehicle into a private field
func (r *Appointment) Vehicle (repo *persistance.Repository) *Vehicle  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    return vehicle
}

func (r *Appointment) Client (repo *persistance.Repository) *Client  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    client := vehicle.Client(repo)
    return client
}