package api

import (
	"net/http"
	"os"

	"github.com/Farky8/FitBond/backEnd/api/handlers"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func SetRouter(ts *handlers.Trainings) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/home/create", func(w http.ResponseWriter, r *http.Request) { ts.HandlePostTraining(w, r) })
	router.HandleFunc("GET /home/search", func(w http.ResponseWriter, r *http.Request) { ts.HandleSearchTraining(w, r) })
	router.HandleFunc("GET /home/search/{id}", func(w http.ResponseWriter, r *http.Request) { ts.HandleGetTraining(w, r) })
	router.HandleFunc("PUT /home/search/{id}",
		func(w http.ResponseWriter, r *http.Request) { ts.HandleTrainingApplication(w, r) })

	testIP := os.Getenv("TEST_IP")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://fit-bond.com",
			"https://fit-bond-front-end-586541250183.europe-central2.run.app",
			testIP,
		},
		AllowedMethods:   []string{"GET", "PUT", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	return n
}
