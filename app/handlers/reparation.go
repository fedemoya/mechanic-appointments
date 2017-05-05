package handlers

import(
    "strconv"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/models"
)

func NewReparation(w http.ResponseWriter, r *http.Request) {

  price := r.FormValue("Price")
  fullPayment := r.FormValue("FullPayment")
  partialPayment := r.FormValue("PartialPayment")
  description := r.FormValue("Description")
  vehicleId := r.FormValue("VehicleId")
  date := r.FormValue("Date")

  log.Println("Received the following Price: " + price)
  log.Println("Received the following FullPayment: " + fullPayment)
  log.Println("Received the following partialPayment: " + partialPayment)
  log.Println("Received the following Description: " + description)
  log.Println("Received the following VehicleId: " + vehicleId)
  log.Println("Received the following Date: " + date)

  parsedPrice, err := strconv.ParseInt(price, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  parsedFullPayment, err := strconv.ParseInt(fullPayment, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  var parsedPartialPayment int64
  if parsedFullPayment == 0 {
    parsedPartialPayment, err = strconv.ParseInt(partialPayment, 10, 64)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }
  }

  parsedVehicleId, err := strconv.ParseInt(vehicleId, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  parsedDate, err := strconv.ParseInt(date, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  // TODO: We need a transaction here in order to make
  // these two insertions atomic

  repository := persistance.NewRepository()
  reparation := &models.Reparation{VehicleId: parsedVehicleId, Date: parsedDate, Description: description, Price: parsedPrice}

  reparation_id, err := repository.Save(reparation)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  var paymentAmount int64

  if parsedFullPayment == 1 {
    paymentAmount = parsedPrice
  } else {
    paymentAmount = parsedPartialPayment
  }

  payment := &models.Payment{ReparationId: reparation_id, Date: parsedDate, Amount: paymentAmount}
  _, err = repository.Save(payment)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  idJson, err := json.Marshal(reparation_id)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  w.Write(idJson)
}

func ReparationDetail(w http.ResponseWriter, r *http.Request) {
  var vars map[string]string = mux.Vars(r)
  id := vars["id"]

  repository := persistance.NewRepository()

  reparation := &models.Reparation{}

  parsedId, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  err = repository.Retrieve(reparation, parsedId)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  payments := []models.Payment{}
  err = repository.Search(&models.Payment{}, &payments, "ReparationId = ?", reparation.Id)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  reparationData := struct {
    Date int64
    Description string
    Price int64
    Payments []models.Payment  
  } {
    reparation.Date,
    reparation.Description,
    reparation.Price,
    payments,
  }

  reparationJson, err := json.Marshal(reparationData)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  w.Write(reparationJson)
}

func ReparationList(w http.ResponseWriter, r *http.Request) {

    value := r.Context().Value("user_id")

    if value == nil {
      errStr := "Missing user_id"
      log.Println(errStr)
      http.Error(w, errStr, http.StatusInternalServerError)
      return
    }

    userId := value.(int64) 

    var vars map[string]string = mux.Vars(r)
    date := vars["date"]

    log.Println("Received the following date: " + date)

    var parsedDate int64 = 0
    var err error 

    if date != "" {
        parsedDate, err = strconv.ParseInt(date, 10, 64)
        if err != nil {
            log.Println(err)
            http.Error(w, err.Error(), 500)
            return
        }
    } else {
      http.Error(w, "Missing date parameter.", 500)
      return
    }

    repository := persistance.NewRepository()

    reparations := []models.Reparation{}

    query := "SELECT reparation.* FROM client, vehicle, reparation "
    query = query + "WHERE client.UserId = ? AND client.Id=vehicle.ClientId "
    query = query + "AND vehicle.Id=reparation.VehicleId AND date(reparation.Date, 'unixepoch')=date(?, 'unixepoch')"
    err = repository.DB.Select(&reparations, query, userId, parsedDate)

    if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
    }

    reparationsData := make([]interface{}, len(reparations))

    for i, reparation := range reparations {

        client := reparation.Client(repository)
        vehicle := reparation.Vehicle(repository)

        reparationData := struct {
            Id int64
            ClientName string
            VehicleDescription string
            Date int64
            Price int64
            Description string
        } {
          reparation.Id,
          client.Name,
          vehicle.Brand + " " + vehicle.Model,
          reparation.Date,
          reparation.Price,
          reparation.Description,
        }

        reparationsData[i] = reparationData
    }

    reparationsJson, err := json.Marshal(reparationsData)
    if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
    }

    w.Write(reparationsJson)
}

