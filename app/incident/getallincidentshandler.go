package incident

import (
	"encoding/json"
	"incidentreport/pkg/response"
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
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  "Error",
				Message: "Unable to get incident",
			},
		)

		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(
		getAllIncidentResponse{
			Status: "success",
			Data:   incidents,
		},
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  "Error",
				Message: "Internal server error",
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
}
