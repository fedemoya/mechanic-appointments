package main

import(
    "log"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "mechanics-backend/app/handlers"
    // "mechanics-backend/app/persistance"
)

func main() {

    init_db()

    r := mux.NewRouter()

    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("html"))))

    r.HandleFunc("/client", handlers.NewClient).Methods("POST")
    r.HandleFunc("/client/{id:[0-9]+}", handlers.ClientDetail).Methods("GET")
    r.HandleFunc("/clients", handlers.ClientList).Methods("GET")
    r.HandleFunc("/clients/debtors", handlers.DebtorList).Methods("GET")

    r.HandleFunc("/appointment", handlers.NewAppointment).Methods("POST")
    r.HandleFunc("/appointment/{id:[0-9]+}", handlers.AppointmentDetail).Methods("GET")
    r.HandleFunc("/appointments", handlers.AppointmentList).Methods("GET")
    r.HandleFunc("/appointments/{date:[0-9]+}", handlers.AppointmentList).Methods("GET")

    r.HandleFunc("/vehicle", handlers.NewVehicle).Methods("POST")

    r.HandleFunc("/reparation", handlers.NewReparation).Methods("POST")
    r.HandleFunc("/reparation/{id:[0-9]+}", handlers.ReparationDetail).Methods("GET")
    r.HandleFunc("/reparations", handlers.ReparationList).Methods("GET")
    r.HandleFunc("/reparations/{date:[0-9]+}", handlers.ReparationList).Methods("GET")

    r.HandleFunc("/payment", handlers.NewPayment).Methods("POST")


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
    // manager := persistance.NewDBSchemaManager("mechanics.db")
    // manager.DropAppTables()
    // manager.CreateAppTables()
}
