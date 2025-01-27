package api

import (
    "net/http"
    "github.com/rs/cors"
    "os"

    "github.com/urfave/negroni"
    "github.com/Farky8/FitBond/backEnd/api/handlers"
)

func SetRouter(ts *handlers.Trainings) http.Handler {
    router := http.NewServeMux()
    router.HandleFunc("/home/create", func(w http.ResponseWriter, r *http.Request) { ts.HandlePostTraining(w, r) })
    router.HandleFunc("GET /home/search", func(w http.ResponseWriter, r *http.Request) { ts.HandleSearchTraining(w, r) })
    router.HandleFunc("/home/search/{id}", func(w http.ResponseWriter, r *http.Request) { ts.HandleGetTraining(w, r) })

    test_ip := os.Getenv("TEST_IP")

    c := cors.New(cors.Options{
	AllowedOrigins: []string{"https://fit-bond.com", "https://fit-bond-front-end-586541250183.europe-central2.run.app", test_ip},
	AllowedMethods: []string{"GET", "POST"},
	AllowedHeaders: []string{"Content-Type", "Authorization"},
	AllowCredentials: true,
    })

    n := negroni.Classic()
    n.Use(c)
    n.UseHandler(router)

    return n
}
