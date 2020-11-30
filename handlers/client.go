package handlers

import(
    "strconv"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "mechanics-backend/persistance"
    "mechanics-backend/models"
)

func NewClient(w http.ResponseWriter, r *http.Request) {

  value := r.Context().Value("user_id")

  if value == nil {
    errStr := "Missing user_id"
    log.Println(errStr)
    http.Error(w, errStr, http.StatusInternalServerError)
    return
  }

  userId := value.(int64)

  name := r.FormValue("Name")

  log.Println("Received the following Name: " + name)

  repository := persistance.NewRepository()

  client := &models.Client{Name: name, UserId: userId}

  id, err := repository.Save(client)
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

func ClientDetail(w http.ResponseWriter, r *http.Request) {
  var vars map[string]string = mux.Vars(r)
  id := vars["id"]

  repository := persistance.NewRepository()

  client := &models.Client{}

  parsedId, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  err = repository.Retrieve(client, parsedId)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  vehicles := []models.Vehicle{}
  err = repository.Search(&models.Vehicle{}, &vehicles, "ClientId = ?", client.Id)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  type VehicleHistory struct {
    Id int64
    VehicleDescription string
    Reparations []models.Reparation
  }

  type ClientDetail struct {
    ClientName string
    VehiclesHistory []VehicleHistory
  }

  clientDetailData := ClientDetail{ClientName: client.Name, VehiclesHistory: nil}

  for _, vehicle := range vehicles {
    reparations := []models.Reparation{}
    err = repository.Search(&models.Reparation{}, &reparations, "VehicleId = ?", vehicle.Id)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }
    vehicleHistory := VehicleHistory{Id: vehicle.Id, VehicleDescription: vehicle.Description(), Reparations: reparations}
    clientDetailData.VehiclesHistory = append(clientDetailData.VehiclesHistory, vehicleHistory)
  }

  clientDetailJson, err := json.Marshal(clientDetailData)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  w.Write(clientDetailJson)
}

func ClientList(w http.ResponseWriter, r *http.Request) {

    value := r.Context().Value("user_id")

    if value == nil {
      errStr := "Missing user_id"
      log.Println(errStr)
      http.Error(w, errStr, http.StatusInternalServerError)
      return
    }

    userId := value.(int64)

    repository := persistance.NewRepository()

    clients := []models.Client{}
    err := repository.Search(&models.Client{}, &clients, "UserId=?", userId)

  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

    clientsJson, err := json.Marshal(clients)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

    w.Write(clientsJson)
}

func DebtorList(w http.ResponseWriter, r *http.Request) {

    value := r.Context().Value("user_id")

    if value == nil {
      errStr := "Missing user_id"
      log.Println(errStr)
      http.Error(w, errStr, http.StatusInternalServerError)
      return
    }

    userId := value.(int64)

    repository := persistance.NewRepository()

    clients := []models.Client{}
    query := "SELECT DISTINCT client.* FROM client, vehicle, reparation, payment "
    query = query + "WHERE client.UserId = ? AND client.Id=vehicle.ClientId "
    query = query + "AND vehicle.Id=reparation.VehicleId "
    query = query + "AND reparation.Id=payment.ReparationId "
    query = query + "GROUP BY client.Id, reparation.Id "
    query = query + "HAVING reparation.Price > SUM(payment.Amount)"
    err := repository.DB.Select(&clients, query, userId)

    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }

    clientsJson, err := json.Marshal(clients)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }

    w.Write(clientsJson)
}
