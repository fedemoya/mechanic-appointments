package persistance

import (
    "strconv"
    _ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
    "database/sql"
    "log"
    "reflect"
)

type Repository struct {
    db *sqlx.DB
}

func NewRepository (dbName string) *Repository  {
    sqlx.NameMapper = func(s string) string { return s }
    // this Pings the database trying to connect, panics on error
    // use sqlx.Open() for sql.Open() semantics
    db, err := sqlx.Connect("sqlite3", "/data/" + dbName)
    if err != nil {
        log.Fatalln(err)
    }
    return &Repository{db}
}

func (r *Repository) Save (object interface{}) (int64, error) {

    v := reflect.ValueOf(object).Elem()
    t := v.Type()

    objectName := t.Name()

    var insert = "INSERT INTO " + objectName + " ("

    insert = insert +  t.Field(0).Name

    for i := 1; i < t.NumField(); i++ {
        insert = insert + ", " +  t.Field(i).Name
    }

    insert  = insert + ") VALUES (null, "

    // skip Field 0 because it is the Id and
    // it is generated by the database
    numField := t.NumField()
    for i := 1; i < numField; i++ {

        field := v.Field(i)
        fieldKind := field.Kind()
        var fieldAsString string
        if (fieldKind.String() == "int64") {
            i := field.Interface().(int64)
            fieldAsString = strconv.FormatInt(i, 10)
        } else if (fieldKind.String() == "int") {
            i := field.Interface().(int)
            fieldAsString = strconv.Itoa(i)
        } else if (fieldKind.String() == "string") {
            fieldAsString = field.Interface().(string)
        } else {
            panic("Unexpected field kind: " + fieldKind.String());
        }

        insert = insert + "\"" + fieldAsString + "\""
        if (i < numField - 1) {
            insert = insert + ", "
        }
    }

    insert = insert + ")"

    var stmt, err = r.db.Prepare(insert)

    var res sql.Result
    res, err = stmt.Exec()

    var id int64
    id, err = res.LastInsertId()

    return id, err
}

func (r *Repository) Retrieve (emptyObject interface{}, id int64) error {

    t := reflect.TypeOf(emptyObject).Elem()

    objectName := t.Name()

    selectStr := "SELECT * FROM " + objectName + " WHERE Id = " + strconv.FormatInt(id, 10)

    err := r.db.Get(emptyObject, selectStr);

    return err
}
