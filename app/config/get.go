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
        log.Fatalln("Undefined Driver Name")
    }

    dataSourceName := os.Getenv("DATA_SOURCE_NAME")

    if dataSourceName == "" {
        log.Fatalln("Undefined Data Source Name")
    }

    configs = make(map[string]string)
    configs["DRIVER_NAME"] = driverName
    configs["DATA_SOURCE_NAME"] = dataSourceName
}

func Get(configName string) (string, error) {

    if configs == nil {
        return "", errors.New("The config store wasn't initialized.")
    }

    return configs[configName], nil
}