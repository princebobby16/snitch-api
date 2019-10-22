package incident

import (
	"incidentreport/db/database"
	"incidentreport/pkg/stringconv"
	"log"
	"strings"
)

func getAllIncidents() (incidents, error) {
	// save directory in database
	var allIncidents incidents
	// save directory in database
	getAllIncidentsStatement := `
		SELECT image_path, "time", location
		FROM incidentreport.incident
		`

	rows, err := database.DBConn.Query(getAllIncidentsStatement)
	if err != nil {
		log.Println(err)
		return allIncidents, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		var location string

		var oneIncident incident

		err = rows.Scan(&oneIncident.Image, &oneIncident.MataData.Time, &location)
		if err != nil {
			log.Println(err)
			allIncidents = append(allIncidents, oneIncident)
			continue
		}

		s := strings.Split(location, ",")

		latitude, err := stringconv.StrtoF(s[0])
		if err != nil {
			log.Println(err)
			return allIncidents, err
		}

		longitude, err := stringconv.StrtoF(s[1])
		if err != nil {
			log.Println(err)
			return allIncidents, err
		}

		oneIncident.MataData.Location.Longitude = longitude
		oneIncident.MataData.Location.Latitude = latitude

		// push to array
		allIncidents = append(allIncidents, oneIncident)
	}

	return allIncidents, nil

}
