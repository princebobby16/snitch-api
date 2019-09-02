package incident

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"incidentreport/pkg/stringconv"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/genproto/googleapis/type/latlng"
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

func HandleAddMetaData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	incidetID := vars["id"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(requestBody)

	var metaData metadataRequest
	metaData.ID, err = stringconv.StrtoI(incidetID)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(requestBody))

	err = json.Unmarshal(requestBody, &metaData)
	if err != nil {
		log.Println(err)
		return
	}

	id, err := AddMetaData(metaData)
	if err != nil {
		log.Println(err)
		return
	}

	_ = json.NewEncoder(w).Encode(
		metadataAddedResponse{
			Status: "success",
			Data: struct {
				ID int `json:"id"`
			}{
				ID: id,
			},
		})
}
