package incident

import (
	"incidentreport/db/database"
	"incidentreport/pkg/stringconv"
	"log"
)

func AddMetaData(metadata metadataRequest) (int, error) {
	// save directory in database
	saveMetaDatatatement := `
		UPDATE incidentreport.incident
		Set location = $1
		WHERE incident_id = $2
		RETURNING incident_id`

	lastInsertedId := 0

	latitude := stringconv.FtoStr(metadata.Location.Latitude)
	longitude := stringconv.FtoStr(metadata.Location.Longitude)

	location := latitude + "," + longitude
	log.Println(metadata.Location)
	log.Println(location)

	err := database.DBConn.QueryRow(saveMetaDatatatement, location, metadata.ID).Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return lastInsertedId, nil
}
