package incident

import (
	"fmt"
	"log"
	"net/http"
)

func HandleImageUpload(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 5 << 20 specifies a maximum
	// upload of 5 MB files.
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Println(err)
		// send error response
	}

	// FormFile returns the first file for the given key `image`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		// send error response
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

	 id,err := SaveIncident(file, handler.Filename)
	if err != nil {
		log.Println(err)
	}

	 log.Println(id)

	_, _ = fmt.Fprintf(w, "%d", id)

}
