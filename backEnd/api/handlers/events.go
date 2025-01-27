package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/Farky8/FitBond/backEnd/api/storage"
)

func (ts *Trainings) HandleGetTraining(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")


    if idStr == "" {
	http.Error(w, "Required query parameter is missing", http.StatusBadRequest)
	return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
	http.Error(w, fmt.Sprintf("Could not convert id to integer: %v", err), http.StatusInternalServerError)
	return
    }

    var queryEvent storage.EventInfo
    
    if err := ts.DB.First(&queryEvent, "id = ?", id); err.Error != nil {
	http.Error(w, fmt.Sprintf("Could not find the film with this ID in the database: %v", err), http.StatusBadRequest)
	return
    }

    payloadJson, err := json.Marshal(queryEvent)
    if err != nil {
	http.Error(w, fmt.Sprintf("Error marshaling the response: %v", err), http.StatusInternalServerError)
	return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(payloadJson)
}

func (ts *Trainings) HandleTrainingApplication(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")


    if idStr == "" {
	http.Error(w, "Required query parameter is missing", http.StatusBadRequest)
	return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
	http.Error(w, fmt.Sprintf("Could not convert id to integer: %v", err), http.StatusInternalServerError)
	return
    }

    var queryEvent storage.EventInfo
    
    if err := ts.DB.First(&queryEvent, "id = ?", id); err.Error != nil {
	http.Error(w, fmt.Sprintf("Could not find the film with this ID in the database: %v", err), http.StatusBadRequest)
	return
    }

    if queryEvent.Applied == queryEvent.Capacity {
	http.Error(w, "The event is full", http.StatusMethodNotAllowed)
	return
    }

    queryEvent.Applied++

    if err := ts.DB.Save(&queryEvent); err.Error != nil {
	http.Error(w, fmt.Sprintf("Failed to save the event to the database: %v", err.Error), http.StatusInternalServerError)
	return
    }

    payloadJson, err := json.Marshal(queryEvent)
    if err != nil {
	http.Error(w, fmt.Sprintf("Error marshaling the response: %v", err), http.StatusInternalServerError)
	return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(payloadJson)
}

