package models

type Appointment struct {
    Id int64
    ClientId int64
    VehicleId int64
    Date int64
}

// func (r *Appointment) Client (repo *persistance.Repository) *models.Client  {
//     vehicle := &Vehicle{}
//     repo.Retrieve(vehicle, r.VehicleId)
//     return vehicle
// }
//
// func (r *Appointment) Vehicle (repo *persistance.Repository) *models.Vehicle  {
//     vehicle := &Vehicle{}
//     repo.Retrieve(vehicle, r.VehicleId)
//     return vehicle
// }
