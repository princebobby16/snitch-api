package incident

import (
	"incidentreport/db/database"
	"incidentreport/pkg/stringconv"
	"log"
	"strings"
)

func getOneIncident(id int64) (incident, error) {

	var oneIncident incident
	// save directory in database
	getOneIncidentStatement := `
		SELECT image_path, location, "time"
		FROM incidentreport.incident
		WHERE incident_id = $1`

	var location string

	row := database.DBConn.QueryRow(getOneIncidentStatement, id)
	err := row.Scan(
		&oneIncident.Image,
		&location,
		&oneIncident.MataData.Time)
	if err != nil {
		log.Println(err)
		return oneIncident, err
	}

	s := strings.Split(location, ",")

	latitude, err := stringconv.StrtoF(s[0])
	if err != nil {
		log.Println(err)
		return oneIncident, err
	}

	longitude, err := stringconv.StrtoF(s[1])
	if err != nil {
		log.Println(err)
		return oneIncident, err
	}

	oneIncident.MataData.Location.Longitude = longitude
	oneIncident.MataData.Location.Latitude = latitude

	return oneIncident, nil
}
