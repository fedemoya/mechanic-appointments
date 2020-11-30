package main

import (
    "database/sql"
    "flag"
    gorillaHandlers "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "log"
    "mechanics-backend/config"
    "mechanics-backend/handlers"
    "mechanics-backend/persistance"
    "net/http"
    "os"
    "time"
)

func main() {

    var initDb = flag.Bool("init-db", false, "Initialize db")
    flag.Parse()

    init_configs()

    if *initDb {
        init_db()
    }

    router := mux.NewRouter()

    apiRouter := router.PathPrefix("/api").Subrouter()

    apiRouter.HandleFunc("/login", handlers.Login).Methods("POST")

    apiRouter.HandleFunc("/client", handlers.NewClient).Methods("POST")
    apiRouter.HandleFunc("/client/{id:[0-9]+}", handlers.ClientDetail).Methods("GET")
    apiRouter.HandleFunc("/clients", handlers.ClientList).Methods("GET")
    apiRouter.HandleFunc("/clients/debtors", handlers.DebtorList).Methods("GET")

    apiRouter.HandleFunc("/appointment", handlers.NewAppointment).Methods("POST")
    apiRouter.HandleFunc("/appointment/{id:[0-9]+}", handlers.AppointmentDetail).Methods("GET")
    apiRouter.HandleFunc("/appointments", handlers.AppointmentList).Methods("GET")
    apiRouter.HandleFunc("/appointments/{date:[0-9]+}", handlers.AppointmentList).Methods("GET")

    apiRouter.HandleFunc("/vehicle", handlers.NewVehicle).Methods("POST")

    apiRouter.HandleFunc("/reparation", handlers.NewReparation).Methods("POST")
    apiRouter.HandleFunc("/reparation/{id:[0-9]+}", handlers.ReparationDetail).Methods("GET")
    apiRouter.HandleFunc("/reparations", handlers.ReparationList).Methods("GET")
    apiRouter.HandleFunc("/reparations/{date:[0-9]+}", handlers.ReparationList).Methods("GET")

    apiRouter.HandleFunc("/payment", handlers.NewPayment).Methods("POST")

    staticFilesDir, err := config.Get("STATIC_FILES_DIR")
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("Serving static files from %s\n", staticFilesDir)
    router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilesDir)))

    routerWithAuth := handlers.NewAuthorizationHandler(router)
    routerWithAuthAndLog := gorillaHandlers.LoggingHandler(os.Stdout, routerWithAuth)

    log.Println("Listening...")

    srv := &http.Server{
        Handler:      routerWithAuthAndLog,
        Addr:         "0.0.0.0:3000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func init_configs() {

    config.Init()
}

func init_db()  {

    log.Println("Initializing database")

    var dbSchemaManager = persistance.NewDBSchemaManager()
    dbSchemaManager.CreateAppTables()

    // Insert user TODO Remove this from here
    driverName, err := config.Get("DRIVER_NAME")
    if err != nil {
        log.Fatalln(err)
    }

    dataSourceName, err := config.Get("DATA_SOURCE_NAME")
    if err != nil {
        log.Fatalln(err)
    }

    db, err := sql.Open(driverName, dataSourceName)
    if err != nil {
        log.Fatalln(err)
    }

    _, err = db.Exec(`INSERT INTO User values (
        NULL,
        'Marcos Basaldella',
        'bure',
        'abc123'
    )`)

    if err != nil {
        log.Fatalln(err)
    }
}
