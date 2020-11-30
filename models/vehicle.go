package models

import(
    "mechanics-backend/persistance"
)

type Vehicle struct {
    Id int64
    ClientId int64
    PlateNumber string
    ChassisNumber string
    Brand string
    Model string
    Year int
}

func (v *Vehicle) Client (repo *persistance.Repository) *Client  {
    client := &Client{}
    repo.Retrieve(client, v.ClientId)
    return client
}

func (v *Vehicle) Description() string {
    return v.Brand + " " + v.Model
}
