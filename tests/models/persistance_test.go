package models

import (
	"testing"
    "os"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/models"
)

func TestSaveClient(t *testing.T) {
    var client *models.Client = &models.Client{Name: "Federico"}
    repository := persistance.NewRepository("mechanics_test.db")
    id, err := repository.Save(client)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Client object")
    }
}

func TestMain(m *testing.M) {
    manager := persistance.NewDBSchemaManager("mechanics_test.db")
    manager.DropAppTables()
    manager.CreateAppTables()
    ret := m.Run()
    manager.DropAppTables()
    os.Exit(ret)
}
