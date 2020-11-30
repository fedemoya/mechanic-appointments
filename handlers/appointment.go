package handlers

import(
    "strconv"
    "log"
    "time"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "mechanics-backend/persistance"
    "mechanics-backend/models"
)

func NewAppointment(w http.ResponseWriter, r *http.Request) {

  vehicleId := r.FormValue("VehicleId")
  date := r.FormValue("Date")

  log.Println("Received the following VehicleId: " + vehicleId)
  log.Println("Received the following Date: " + date)

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

  repository := persistance.NewRepository()
  appointment := &models.Appointment{VehicleId: parsedVehicleId, Date: parsedDate}

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

  repository := persistance.NewRepository()

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
    }

    repository := persistance.NewRepository()

    appointments := []models.Appointment{}
    
    if parsedDate == 0 {
        parsedDate = time.Now().Unix()
    }

    query := "SELECT appointment.* FROM client, vehicle, appointment "
    query = query + "WHERE client.UserId = ? AND client.Id=vehicle.ClientId "
    query = query + "AND vehicle.Id=appointment.VehicleId AND date(appointment.Date, 'unixepoch')=date(?, 'unixepoch')"
    err = repository.DB.Select(&appointments, query, userId, parsedDate)

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
