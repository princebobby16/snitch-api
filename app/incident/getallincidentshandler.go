package incident

import (
	"encoding/json"
	"log"
	"net/http"
)

type incidents []incident

type getAllIncidentResponse struct {
	Status string    `json:"status"`
	Data   incidents `json:"data"`
}

func HandleGetAllIncidents(w http.ResponseWriter, _ *http.Request) {

	incidents, err := getAllIncidents()
	if err != nil {
		log.Println(err)
		return
	}

	_ = json.NewEncoder(w).Encode(
		getAllIncidentResponse{
			Status: "succes",
			Data:   incidents,
		})
}
