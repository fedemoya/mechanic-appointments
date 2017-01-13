package models

import(
    "mechanics-backend/app/persistance"
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

func (r *Vehicle) Client (repo *persistance.Repository) *Client  {
    client := &Client{}
    repo.Retrieve(client, r.ClientId)
    return client
}
