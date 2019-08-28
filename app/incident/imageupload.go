package incident

import (
	"bytes"
	"incidentreport/db/database"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

func SaveIncidentImage(file multipart.File, filename string) (int, error) {
	filepath := "/home/naru/Pictures/snitch/" + filename
	f, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println(err)
		return 0, err
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// save directory in database
	saveDirectoryStatement := `
		INSERT INTO incidentreport.incident (image_path, "time")
		VALUES ($1, $2)
		RETURNING incident_id`

	lastInsertedId := 0

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	err = database.DBConn.QueryRow(saveDirectoryStatement, filepath, now).Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return lastInsertedId, nil
}
