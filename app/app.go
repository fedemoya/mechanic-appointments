package main

import(
    "log"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "mechanics-backend/app/handlers"
)

func main() {

    r := mux.NewRouter()
    r.Handle("/client", &handlers.ClientHandler{})

    log.Println("Listening...")

    srv := &http.Server{
        Handler:      r,
        Addr:         "0.0.0.0:3000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func init_db()  {
    
}
