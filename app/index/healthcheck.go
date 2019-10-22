package index

import (
	"encoding/json"
	"fmt"
	"incidentreport/pkg/response"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(
		Alive{
			Alive:       true,
			Author:      "Benjy Asiamah-Koranteng",
			Version:     "0.0.1",
			Environment: "development",
		},
	)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(
			response.ErrorResponse{
				Status:  fmt.Sprintf("%s", http.StatusInternalServerError),
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
