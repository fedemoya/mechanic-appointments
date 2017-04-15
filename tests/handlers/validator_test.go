package handlers

import (
    "testing"
    "mechanics-backend/app/handlers"
)

func TestRequiredName(t *testing.T) {
    var name string
    var err error
    name, err = handlers.Validate("Jose Mercado").Required().AsName()
    if err != nil {
        t.Error(err)
    }
    if name != "Jose Mercado" {
        t.Error("Incorrect name returned")
    }
    _, err = handlers.Validate("").Required().AsName()
    if err.Error() != "El parametro no puede ser nulo." {
        t.Error("Wrong error returned")
    }
}

func TestOptionalName(t *testing.T) {
    var err error
    _, err = handlers.Validate("").AsName()
    if err != nil {
        t.Error(err)
    }
    var name string
    name, err = handlers.Validate("Antonio Moya").AsName()
    if err != nil {
        t.Error(err)
    }
    if name != "Antonio Moya" {
        t.Error("Incorrect name returned")
    }
}

func TestBadName(t *testing.T) {
    var err error
    _, err = handlers.Validate("Antonio 1234 Moya").AsName()
    if err == nil {
        t.Error("Validation failed")
    }
    if err.Error() != "Nombre invalido." {
        t.Error("Wrong error")
    }
}

func TestRequiredPrice(t *testing.T) {
    var price float64
    var err error
    price, err = handlers.Validate("100").Required().AsPrice()
    if err != nil {
        t.Error(err)
    }
    if price != 100 {
        t.Error("Incorrect price returned")
    }
}