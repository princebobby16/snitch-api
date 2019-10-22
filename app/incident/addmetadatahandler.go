package incident

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/genproto/googleapis/type/latlng"
	"incidentreport/pkg/response"
	"incidentreport/pkg/stringconv"
	"io/ioutil"
	"log"
	"net/http"
)

type metadataRequest struct {
	Location latlng.LatLng `json:"location"`
	Time     string        `json:"time"`
	ID       int64         `json:"id"`
}

type metadataAddedResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID int `json:"id"`
	} `json:"data"`
}

type badRequestResponse struct {
	Status string `json:"status"`
	Data   struct {
		Location string `json:"location"`
		Time     string `json:"time"`
		ID       string `json:"id"`
	} `json:"data"`
}

func HandleAddMetaData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	incidetID := vars["id"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  "fail",
				Message: "Bad request",
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	log.Println(requestBody)

	var metaData metadataRequest
	metaData.ID, err = stringconv.StrtoI(incidetID)
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

	log.Println(string(requestBody))

	err = json.Unmarshal(requestBody, &metaData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(
			badRequestResponse{
				Status: "Error",
				Data: struct {
					Location string `json:"location"`
					Time     string `json:"time"`
					ID       string `json:"id"`
				}{
					Location: "Invalid location",
					Time:     "Invalid time",
					ID:       "Invalid ID",
				},
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	id, err := AddMetaData(metaData)
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

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(
		metadataAddedResponse{
			Status: "success",
			Data: struct {
				ID int `json:"id"`
			}{
				ID: id,
			},
		})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  "fail",
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
