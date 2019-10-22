package incident

import (
	"encoding/json"
	"google.golang.org/genproto/googleapis/type/latlng"
	"incidentreport/pkg/response"
	"incidentreport/pkg/stringconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type incident struct {
	Image    string `json:"image"`
	MataData struct {
		Location latlng.LatLng `json:"location"`
		Time     string        `json:"time"`
	} `json:"mata_data"`
}

type getOneIncidentResponse struct {
	Status string   `json:"status"`
	Data   incident `json:"data"`
}

type getIdFailResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID string `json:"id"`
	} `json:"data"`
}

func HandleGetOneIncident(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	incidetID := vars["id"]

	id, err := stringconv.StrtoI(incidetID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Error",
			Message: "Unable to complete request",
		})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	oneIncident, err := getOneIncident(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(getIdFailResponse{
			Status: "Fail",
			Data: struct {
				ID string `json:"id"`
			}{
				ID: "Invalid ID",
			},
		})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(getOneIncidentResponse{
		Status: "success",
		Data:   oneIncident,
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  "Error",
				Message: "Unable to complete request",
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
}
