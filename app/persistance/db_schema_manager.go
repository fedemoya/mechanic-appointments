
package persistance

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "log"
)

type DBSchemaManager struct {
    db *sql.DB
}

func NewDBSchemaManager (dbName string) *DBSchemaManager  {
    db, err := sql.Open("sqlite3", "/data/" + dbName)
    if err != nil {
        log.Fatalln(err)
    }
    return &DBSchemaManager{db}
}

func (manager *DBSchemaManager) CreateAppTables()  {

    _, err := manager.db.Exec(`CREATE TABLE Client (
        Id integer primary key,
        Name text
    )`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`CREATE TABLE Vehicle (
        Id integer primary key,
        ClientId integer,
        PlateNumber text,
        ChassisNumber text,
        Brand text,
        Model text,
        Year integer
    )`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`CREATE TABLE Appointment (
        Id integer primary key,
        VehicleId int64,
        Date integer
    )`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`CREATE TABLE Reparation (
        Id integer primary key,
        VehicleId int64,
        Description string,
        Date integer
        Price int64
    )`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`CREATE TABLE Payment (
        Id integer primary key,
        ReparationId int64,
        Date int64,
        Amount int64
    )`)

    if err != nil {
        log.Fatalln(err)
    }

}

func (manager *DBSchemaManager) DropAppTables()  {

    _, err := manager.db.Exec(`DROP TABLE IF EXISTS Client`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`DROP TABLE IF EXISTS Vehicle`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`DROP TABLE IF EXISTS Appointment`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`DROP TABLE IF EXISTS Reparation`)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = manager.db.Exec(`DROP TABLE IF EXISTS Payment`)

    if err != nil {
        log.Fatalln(err)
    }

}

func (manager *DBSchemaManager) CreateTestTables()  {

    _, err := manager.db.Exec(`
        CREATE TABLE Person
        (Id integer primary key,
        FirstName text,
        LastName text, Email text,
        BirthDate integer)
    `)

    if err != nil {
        log.Fatalln("Cannot create Person table: " + err.Error())
    }
}

func (manager *DBSchemaManager) DropTestTables()  {
    _, err := manager.db.Exec("DROP TABLE IF EXISTS Person")

    if err != nil {
        log.Fatalln("Cannot drop Person table: " + err.Error())
    }
}
