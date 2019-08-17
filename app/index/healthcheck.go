package index

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(
		Alive{
			Alive:       true,
			Author:      "Benjy Asiamah-Koranteng",
			Version:     "0.0.1",
			Environment: "development",
		},
	)

}
