package persistance

import (
	"testing"
    "os"
    "time"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/config"
)

type Person struct {
    Id        int64
    FirstName string
    LastName  string
    Email     string
    BirthDate int64
}

func TestSave(t *testing.T) {

    repository := persistance.NewRepository()
    
    var birthDate time.Time = time.Date(1986, time.May, 24, 7, 0, 0, 0, time.UTC)
    p := &Person{FirstName: "Federico", LastName: "Moya", Email: "federicoamoya@gmail.com", BirthDate: birthDate.Unix()}
    id, err := repository.Save(p)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Person object")
    }

}

func TestRetrieve(t *testing.T)  {
    
    repository := persistance.NewRepository()
    
    var birthDate time.Time = time.Date(1986, time.May, 24, 7, 0, 0, 0, time.UTC)
    p := &Person{FirstName: "Federico", LastName: "Moya", Email: "federicoamoya@gmail.com", BirthDate: birthDate.Unix()}
    id, err := repository.Save(p)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Person object")
    }

    var pAgain *Person = &Person{}
    err = repository.Retrieve(pAgain, id)

    if err != nil {
        t.Error(err)
    }

    if pAgain.FirstName != "Federico" {
        t.Error("Problem occur retrieving Person object")
    }
}

func TestSearch(t *testing.T)  {
    
    repository := persistance.NewRepository()

    var birthDate time.Time = time.Date(1986, time.May, 24, 7, 0, 0, 0, time.UTC)
    p := &Person{FirstName: "Federico", LastName: "Moya", Email: "federicoamoya@gmail.com", BirthDate: birthDate.Unix()}
    id, err := repository.Save(p)

    if err != nil {
        t.Error(err)
    }

    if (id == 0) {
        t.Error("Problem occur saving Person object")
    }

    persons := []Person{}
    err = repository.Search(&Person{}, &persons, "")

    if err != nil {
        t.Error(err)
    }

    if persons[0].FirstName != "Federico" {
        t.Error("Problem occur retrieving Person object")
    }
}

func TestMain(m *testing.M) {

    os.Setenv("DRIVER_NAME", "sqlite3")
    os.Setenv("DATA_SOURCE_NAME", "/data/mechanics_test.db")

    config.Init()

    manager := persistance.NewDBSchemaManager()
    
    manager.DropTestTables()
    manager.CreateTestTables()
    ret := m.Run()
    manager.DropTestTables()
    os.Exit(ret)
}
