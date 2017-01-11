
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

// func (manager *DBSchemaManager) CreateAppTables()  {
//
//     _, err := manager.db.Exec("CREATE TABLE Client (Id integer primary key,
//         Name text)");
//
//     log.Println(err);
//
//     _, err := manager.db.Exec("CREATE TABLE Vehicle (Id integer primary key,
//         ClientId integer,
//         PlateNumber text,
//         ChassisNumber text,
//         Brand text,
//         Model text
//         Year integer)");
//
//     log.Println(err);
//
//     _, err := manager.db.Exec("CREATE TABLE Appointment (Id integer primary key,
//         ClientId int64
//         VehicleId int64)");
//
//     log.Println(err);
//
//     _, err := manager.db.Exec("CREATE TABLE Reparation (Id integer primary key,
//         VehicleId int64
//         Description string)");
//
//     log.Println(err);
// }

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
