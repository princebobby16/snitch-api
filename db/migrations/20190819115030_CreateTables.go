
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20190819115030(txn *sql.Tx) error {
	_, err := txn.Exec("CREATE TABLE IF NOT EXISTS incidentReport.incident(" +
		"incident_id SERIAL PRIMARY KEY," +
		"image_path VARCHAR(250)," +
		"location VARCHAR(250)," +
		"time TIMESTAMP," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP) ;")
	if err != nil {
		return err
	}
	return nil
}

// Down is executed when this migration is rolled back
func Down_20190819115030(txn *sql.Tx) error {
	_, err := txn.Exec("DROP TABLE incidentReport.incident CASCADE ;")
	if err != nil {
		return err
	}
	return nil
}
