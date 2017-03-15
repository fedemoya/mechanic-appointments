package models

import (
    "mechanics-backend/app/persistance"
)

type Reparation struct {
    Id int64
    VehicleId int64
    Description string
    Date int64
    Price int64
}

func (r *Reparation) Vehicle (repo *persistance.Repository) *Vehicle  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    return vehicle
}

func (r *Reparation) Client (repo *persistance.Repository) *Client  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    client := vehicle.Client(repo)
    return client
}