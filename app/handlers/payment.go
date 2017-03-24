package handlers

import(
    "strconv"
    "log"
    "net/http"
    "encoding/json"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/models"
)

func NewPayment(w http.ResponseWriter, r *http.Request) {

  reparationId := r.FormValue("ReparationId")
  date := r.FormValue("Date")
  amount := r.FormValue("Amount")

  log.Println("Received the following ReparationId: " + reparationId)
  log.Println("Received the following Date: " + date)
  log.Println("Received the following Amount: " + amount)

  parsedReparationId, err := strconv.ParseInt(reparationId, 10, 64)
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

  parsedAmount, err := strconv.ParseInt(amount, 10, 64)
  if err != nil {
      log.Println(err)
      http.Error(w, err.Error(), 500)
      return
  }

  repository := persistance.NewRepository("mechanics.db")

  payment := &models.Payment{ReparationId: parsedReparationId, Date: parsedDate, Amount: parsedAmount}

  id, err := repository.Save(payment)
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
