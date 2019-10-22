package incident

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type incidentCreatedResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID int `json:"id"`
	} `json:"data"`
	Links struct {
		Rel    string `json:"rel"`
		Href   string `json:"href"`
		Action string `json:"action"`
	} `json:"links"`
}

type failResponse struct {
	Status string `json:"status"`
	Data   struct {
		Image string `json:"image"`
	} `json:"data"`
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"data"`
}

func HandleImageUpload(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 5 << 20 specifies a maximum
	// upload of 5 MB files.
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			failResponse{
				Status: "fail",
				Data: struct {
					Image string `json:"image"`
				}{
					Image: "file size too large",
				},
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	// FormFile returns the first file for the given key `image`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(
			errorResponse{
				Status:  "error",
				Message: "could not read file",
			})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	id, err := SaveIncidentImage(file, handler.Filename)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(
			errorResponse{
				Status:  "error",
				Message: "could not create incident",
			},
		)

		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(
		incidentCreatedResponse{
			Status: "success",
			Data: struct {
				ID int `json:"id"`
			}{
				ID: id,
			},
			Links: struct {
				Rel    string `json:"rel"`
				Href   string `json:"href"`
				Action string `json:"action"`
			}{
				Rel:    "meta data",
				Href:   "",
				Action: "PUT",
			},
		})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(
			errorResponse{
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
