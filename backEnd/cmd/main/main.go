package main

import (
    "net/http"
    "log"
    "gorm.io/gorm"

    "github.com/Farky8/FitBond/backEnd/api"
    "github.com/Farky8/FitBond/backEnd/api/storage"
    "github.com/Farky8/FitBond/backEnd/api/handlers"
)

func main() {
    //var trainings *storage.Trainings = storage.DBSetUp()
    var trainings *gorm.DB = storage.DBSetUp()
    
    var router http.Handler = api.SetRouter(&handlers.Trainings{
	DB: trainings,
    })
    log.Fatal(http.ListenAndServe(":8080", router))
}
