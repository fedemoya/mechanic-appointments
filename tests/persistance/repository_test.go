package persistance

import (
	"testing"
    "os"
    "time"
    "mechanics-backend/app/persistance"
)

type Person struct {
    Id        int64
    FirstName string
    LastName  string
    Email     string
    BirthDate int64
}

func TestSave(t *testing.T) {

    repository := persistance.NewRepository("mechanics_test.db")
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
    repository := persistance.NewRepository("mechanics_test.db")
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

func TestMain(m *testing.M) {
    manager := persistance.NewDBSchemaManager("mechanics_test.db")
    manager.DropTestTables()
    manager.CreateTestTables()
    ret := m.Run()
    manager.DropTestTables()
    os.Exit(ret)
}
