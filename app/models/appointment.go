package models

import (
    "mechanics-backend/app/persistance"
)

type Appointment struct {
    Id int64
    ClientId int64
    VehicleId int64
    Date int64
}

func (r *Appointment) Client (repo *persistance.Repository) *Client  {
    client := &Client{}
    repo.Retrieve(client, r.ClientId)
    return client
}

func (r *Appointment) Vehicle (repo *persistance.Repository) *Vehicle  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    return vehicle
}
