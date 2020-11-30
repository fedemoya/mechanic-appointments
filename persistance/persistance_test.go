package persistance

import (
    "testing"
    "os"
    "time"
    "mechanics-backend/models"
    "mechanics-backend/config"
)

func TestSaveRetrieveClient(t *testing.T) {
    
    var client *models.Client = &models.Client{Name: "Federico"}
    
    
    repository := NewRepository()
    
    id, err := repository.Save(client)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Client object")
    }

    clientAgain := &models.Client{}
    err = repository.Retrieve(clientAgain, id)

    if err != nil {
        t.Error(err)
    }

    if clientAgain.Name != "Federico" {
        t.Error("Problem occur retrieving Client object")
    }
}

func TestSaveRetrieveAppointment(t *testing.T) {
    
    var appointmentDate time.Time = time.Date(2017, time.May, 24, 7, 0, 0, 0, time.UTC)
    var appointment *models.Appointment = &models.Appointment{VehicleId: 1, Date: appointmentDate.Unix()}
    
    repository := NewRepository()
    
    id, err := repository.Save(appointment)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Appointment object")
    }

    appointmentAgain := &models.Appointment{}
    err = repository.Retrieve(appointmentAgain, id)

    if err != nil {
        t.Error(err)
    }
    
    if appointmentAgain.VehicleId != 1 {
        t.Error("Problem occur retrieving Appointment object")
    }

}

func TestSaveRetrieveReparation(t *testing.T) {
    
    var reparation *models.Reparation = &models.Reparation{VehicleId: 1, Description: "Cambio de aceite y filtros."}
    
    repository := NewRepository()
    
    id, err := repository.Save(reparation)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Reparation object")
    }

    reparationAgain := &models.Reparation{}
    err = repository.Retrieve(reparationAgain, id)

    if err != nil {
        t.Error(err)
    }

    if reparationAgain.VehicleId != 1 {
        t.Error("Problem occur retrieving Reparation object")
    }
}

func TestSaveRetrieveVehicle(t *testing.T) {
    
    var vehicle *models.Vehicle = &models.Vehicle{
        ClientId: 1,
        PlateNumber: "IYN751",
        ChassisNumber: "D8DUD8DYDGNVH764",
        Brand: "Citroen",
        Model: "C4",
        Year: 2010,
    }
    
    repository := NewRepository()
    
    id, err := repository.Save(vehicle)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Vehicle object")
    }

    vehicleAgain := &models.Vehicle{}
    err = repository.Retrieve(vehicleAgain, id)

    if err != nil {
        t.Error(err)
    }

    if vehicleAgain.PlateNumber != "IYN751" {
        t.Error("Problem occur retrieving Vehicle object")
    }
}

func TestVehicleRelations(t *testing.T)  {

    var client *models.Client = &models.Client{Name: "Federico"}
    
    repository := NewRepository()
    
    id, err := repository.Save(client)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Client object")
    }

    var vehicle *models.Vehicle = &models.Vehicle{
        ClientId: id,
        PlateNumber: "IYN751",
        ChassisNumber: "D8DUD8DYDGNVH764",
        Brand: "Citroen",
        Model: "C4",
        Year: 2010,
    }

    id, err = repository.Save(vehicle)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Vehicle object")
    }

    vehicleAgain := &models.Vehicle{}
    err = repository.Retrieve(vehicleAgain, id)

    if err != nil {
        t.Error(err)
    }

    clientAgain := vehicleAgain.Client(repository)

    if clientAgain.Name != "Federico" {
        t.Error("Problem with Vehicle -> Client relation")
    }
}

func TestMain(m *testing.M) {
    
    os.Setenv("DRIVER_NAME", "sqlite3")
    os.Setenv("DATA_SOURCE_NAME", "/data/mechanics_test.db")

    config.Init()

    manager := NewDBSchemaManager()
    manager.DropAppTables()
    manager.CreateAppTables()
    ret := m.Run()
    manager.DropAppTables()
    os.Exit(ret)
}
