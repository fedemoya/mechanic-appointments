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

    s := r.MatcherFunc(func (r *http.Request, match *mux.RouteMatch) bool {
        log.Println("MatcherFunc called.")
        return true
    }).Subrouter()

    s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("html"))))

    s.HandleFunc("/client", handlers.NewClient).Methods("POST")
    s.HandleFunc("/client/{id:[0-9]+}", handlers.ClientDetail).Methods("GET")
    s.HandleFunc("/clients", handlers.ClientList).Methods("GET")
    s.HandleFunc("/clients/debtors", handlers.DebtorList).Methods("GET")

    s.HandleFunc("/appointment", handlers.NewAppointment).Methods("POST")
    s.HandleFunc("/appointment/{id:[0-9]+}", handlers.AppointmentDetail).Methods("GET")
    s.HandleFunc("/appointments", handlers.AppointmentList).Methods("GET")
    s.HandleFunc("/appointments/{date:[0-9]+}", handlers.AppointmentList).Methods("GET")

    s.HandleFunc("/vehicle", handlers.NewVehicle).Methods("POST")

    s.HandleFunc("/reparation", handlers.NewReparation).Methods("POST")
    s.HandleFunc("/reparation/{id:[0-9]+}", handlers.ReparationDetail).Methods("GET")
    s.HandleFunc("/reparations", handlers.ReparationList).Methods("GET")
    s.HandleFunc("/reparations/{date:[0-9]+}", handlers.ReparationList).Methods("GET")

    s.HandleFunc("/payment", handlers.NewPayment).Methods("POST")


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
