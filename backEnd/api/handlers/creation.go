package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/Farky8/FitBond/backEnd/api/storage"
)

func (ts *Trainings) HandlePostTraining(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodOptions {
	w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.WriteHeader(http.StatusOK)
        return
    }

    var newEvent storage.EventInfo
    decode := json.NewDecoder(r.Body)

    if err := decode.Decode(&newEvent); err != nil {
	http.Error(w, fmt.Sprintf("Failed to decode Json: %v", err), http.StatusInternalServerError)
	return
    }

    // ID generated automatically by gorm and Applied 0 by default

    if err := ts.DB.Create(&newEvent); err.Error != nil {
	http.Error(w, fmt.Sprintf("Error inserting to the database %v", err.Error), http.StatusInternalServerError)
	return
    }

    w.WriteHeader(http.StatusCreated)
}
