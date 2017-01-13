package models

// import (
//     "mechanics-backend/app/persistance"
//     "mechanics-backend/app/models"
// )

type Reparation struct {
    Id int64
    VehicleId int64
    Description string
}

// func (r *Reparation) Vehicle (repo *persistance.Repository) *models.Vehicle  {
//     vehicle := &Vehicle{}
//     repo.Retrieve(vehicle, r.VehicleId)
//     return vehicle
// }
