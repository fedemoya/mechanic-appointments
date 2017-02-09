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

func NewAppointment(w http.ResponseWriter, r *http.Request) {

  clientId := r.FormValue("ClientId")
  vehicleId := r.FormValue("VehicleId")
  date := r.FormValue("Date")

  log.Println("Received the following ClientId: " + clientId)
  log.Println("Received the following VehicleId: " + vehicleId)
  log.Println("Received the following Date: " + date)


  repository := persistance.NewRepository("mechanics.db")

  parsedClientId, err := strconv.ParseInt(clientId, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
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

  appointment := &models.Appointment{ClientId: parsedClientId, VehicleId: parsedVehicleId, Date: parsedDate}

  id, err := repository.Save(appointment)
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

func AppointmentDetail(w http.ResponseWriter, r *http.Request) {
  var vars map[string]string = mux.Vars(r)
  id := vars["id"]

  repository := persistance.NewRepository("mechanics.db")

  appointment := &models.Appointment{}

  parsedId, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  err = repository.Retrieve(appointment, parsedId)

  if err != nil {
      log.Println(err)
      appointmentJson, err := json.Marshal(nil)
      if err != nil {
          log.Println(err)
          http.Error(w, err.Error(), 500)
          return
      }
      w.Write(appointmentJson)
      return
  }

  client := appointment.Client(repository)
  vehicle := appointment.Vehicle(repository)

  appointmentData := struct {
      AppointmentId int64
      ClientName string
      VehicleDescription string
      Date int64
  } {
    appointment.Id,
    client.Name,
    vehicle.Brand + " " + vehicle.Model,
    appointment.Date,
  }

  appointmentJson, err := json.Marshal(appointmentData)
  if err != nil {
    log.Println(err)
    http.Error(w, err.Error(), 500)
    return
  }

  w.Write(appointmentJson)
}

func AppointmentList(w http.ResponseWriter, r *http.Request) {

    repository := persistance.NewRepository("mechanics.db")

    appointments := []models.Appointment{}
    err := repository.Search(&models.Appointment{}, &appointments)

    if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
    }

    appointmentsData := make([]interface{}, len(appointments))

    for i, appointment := range appointments {

        client := appointment.Client(repository)
        vehicle := appointment.Vehicle(repository)

        appointmentData := struct {
            AppointmentId int64
            ClientName string
            VehicleDescription string
            Date int64
        } {
          appointment.Id,
          client.Name,
          vehicle.Brand + " " + vehicle.Model,
          appointment.Date,
        }

        appointmentsData[i] = appointmentData
    }

    appointmentsJson, err := json.Marshal(appointmentsData)
    if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
    }

    w.Write(appointmentsJson)
}
