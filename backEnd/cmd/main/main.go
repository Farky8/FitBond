package main

import (
    "net/http"

    "github.com/Farky8/FitBond/backEnd/api"
)

func main() {
    var trainings *storage.Trainings = storage.DBSetUp()
    
    var router http.Handler = api.SetRouter(trainings)
    log.Fatal(http.ListenAndServe(":8080", router))
}
