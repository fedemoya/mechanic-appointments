package config

import (
    "os"
    "errors"
    "log"
)

var configs map[string]string

func Init() {

    driverName := os.Getenv("DRIVER_NAME")

    if driverName == "" {
        driverName = "sqlite3"
    }

    dataSourceName := os.Getenv("DATA_SOURCE_NAME")

    if dataSourceName == "" {
        log.Fatalln("Undefined Data Source Name")
    }

    staticFilesDir := os.Getenv("STATIC_FILES_DIR")

    if staticFilesDir == "" {
        log.Fatalln("Undefined Static files directory")
    }

    configs = make(map[string]string)
    configs["DRIVER_NAME"] = driverName
    configs["DATA_SOURCE_NAME"] = dataSourceName
    configs["STATIC_FILES_DIR"] = staticFilesDir
}

func Get(configName string) (string, error) {

    if configs == nil {
        return "", errors.New("The config store wasn't initialized.")
    }

    return configs[configName], nil
}