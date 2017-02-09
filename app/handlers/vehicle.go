package handlers

import(
    "strconv"
    "log"
    "net/http"
    "encoding/json"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/models"
)

func NewVehicle(w http.ResponseWriter, r *http.Request) {

  clientId := r.FormValue("ClientId")
  plateNumber := r.FormValue("PlateNumber")
  chassisNumber := r.FormValue("ChassisNumber")
  brand := r.FormValue("Brand")
  model := r.FormValue("Model")
  year := r.FormValue("Year")

  log.Println("Received the following ClientId: " + clientId)
  log.Println("Received the following PlateNumber: " + plateNumber)
  log.Println("Received the following ChassisNumber: " + chassisNumber)
  log.Println("Received the following Brand: " + brand)
  log.Println("Received the following Model: " + model)
  log.Println("Received the following Year: " + year)

  repository := persistance.NewRepository("mechanics.db")

  parsedClientId, err := strconv.ParseInt(clientId, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  parsedYear, err := strconv.ParseInt(year, 10, 32)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  vehicle := &models.Vehicle{ClientId: parsedClientId, PlateNumber: plateNumber, ChassisNumber: chassisNumber, Brand: brand, Model: model, Year: int(parsedYear)}

  id, err := repository.Save(vehicle)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  idJson, err := json.Marshal(id)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  w.Write(idJson)
}

// func ClientDetail(w http.ResponseWriter, r *http.Request) {
//   var vars map[string]string = mux.Vars(r)
//   id := vars["id"]
//
//   repository := persistance.NewRepository("mechanics.db")
//
//   client := &models.Client{}
//
//   parsedId, err := strconv.ParseInt(id, 10, 64)
//   if err != nil {
//       log.Println(err)
//   }
//
//   err = repository.Retrieve(client, parsedId)
//   if err != nil {
//       log.Println(err)
//       client = nil
//   }
//
//   clientJson, err := json.Marshal(client)
//   if err != nil {
//       log.Println(err)
//   }
//
//   w.Write(clientJson)
// }
//
// func ClientList(w http.ResponseWriter, r *http.Request) {
//
//     repository := persistance.NewRepository("mechanics.db")
//
//     clients := []models.Client{}
//     err := repository.Search(&models.Client{}, &clients)
//
//     if err != nil {
//         log.Println(err)
//     }
//
//     clientsJson, err := json.Marshal(clients)
//     if err != nil {
//         log.Println(err)
//     }
//
//     w.Write(clientsJson)
// }
