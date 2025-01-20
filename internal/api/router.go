package api

import (
    "net/http"

    "github.com/Farky8/TrainMe/internal/api/storage"
)

func SetRouter(ts *storage.Trainings) http.Handler {
    router := http.NewServeMux()
    router.HandleFunc("POST /home/create", func(w http.ResponseWriter, r *http.Request) { ts.HandlePostTraining(w, r) })
    router.HandleFunc("GET /home/search", func(w http.ResponseWriter, r *http.Request) { ts.HandleSearchTraining(w, r) })
    router.HandleFunc("POST /home/search/{id}", func(w http.ResponseWriter, r *http.Request) { ts.HandleGetTraining(w, r) })
}
