package handlers

import (
    "net/http"
    "fmt"
    "strconv"
    "encoding/json"

    "github.com/Farky8/FitBond/backEnd/api/storage"
)

type SearchResponse struct {
    Length  int			`json:"length"`
    Results []storage.EventInfo	`json:"results"`
}

func (ts *Trainings) HandleSearchTraining(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.Query()

    city := queryParams.Get("city")
    indexStr := queryParams.Get("start")
    lengthStr := queryParams.Get("length")

    if city == "" || indexStr == "" || lengthStr == "" {
	http.Error(w, "Required query parameter is missing", http.StatusBadRequest)
	return
    }

    index, err := strconv.Atoi(indexStr)
    if err != nil {
	http.Error(w, fmt.Sprintf("Could not convert start to integer: %v", err), http.StatusInternalServerError)
	return
    }
    length, err := strconv.Atoi(lengthStr)
    if err != nil {
	http.Error(w, fmt.Sprintf("Could not convert length to integer: %v", err), http.StatusInternalServerError)
	return
    }

    var events []storage.EventInfo
    if err := ts.DB.Where("city = ?", city).Offset(index).Limit(length).Find(&events); err.Error != nil {
	http.Error(w, fmt.Sprintf("Searching in database failed: %v", err.Error), http.StatusBadRequest)
	return
    }

    response := SearchResponse{
	Length: len(events),
	Results: events,
    }

    payloadJson, err := json.Marshal(response)
    if err != nil {
	http.Error(w, fmt.Sprintf("Failed to marshal the response %v", err), http.StatusInternalServerError)
	return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(payloadJson)

}
