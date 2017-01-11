package handlers

import(
    "log"
    "net/http"
    "encoding/json"
    "mechanics-backend/app/models"
)

type ClientHandler struct {
}

func (th *ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  client := &models.Client{Id: 1234, Name: "Federico Moya"}
  clientJson, err := json.Marshal(client)
  if err != nil {
      log.Fatal(err)
      return
  }
  w.Write(clientJson)
}
