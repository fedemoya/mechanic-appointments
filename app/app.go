package main

import(
    "log"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "mechanics-backend/app/persistance"
    "mechanics-backend/app/handlers"
)

func main() {

    init_db()

    r := mux.NewRouter()

    r.HandleFunc("/client", handlers.NewClient).Methods("POST")
    r.HandleFunc("/client/{id:[0-9]+}", handlers.ClientDetail).Methods("GET")

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
    manager := persistance.NewDBSchemaManager("mechanics.db")
    manager.DropAppTables()
    manager.CreateAppTables()
}
