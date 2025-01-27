package main

import (
	"log"
	"net/http"

	"github.com/Farky8/FitBond/backEnd/api"
	"github.com/Farky8/FitBond/backEnd/api/handlers"
	"github.com/Farky8/FitBond/backEnd/api/storage"
)

func main() {
	trainings := storage.DBSetUp()

	router := api.SetRouter(&handlers.Trainings{
		DB: trainings,
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
