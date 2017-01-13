package models

import (
    "mechanics-backend/app/persistance"
)

type Reparation struct {
    Id int64
    VehicleId int64
    Description string
}

func (r *Reparation) Vehicle (repo *persistance.Repository) *Vehicle  {
    vehicle := &Vehicle{}
    repo.Retrieve(vehicle, r.VehicleId)
    return vehicle
}
