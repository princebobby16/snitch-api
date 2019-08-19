
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20190819113843(txn *sql.Tx) error {
	_, err := txn.Exec("CREATE SCHEMA IF NOT EXISTS incidentReport AUTHORIZATION snitch;")
	if err != nil {
		return err
	}
	return nil
}

// Down is executed when this migration is rolled back
func Down_20190819113843(txn *sql.Tx) error {
	_, err := txn.Exec("DROP SCHEMA IF EXISTS incidentReport CASCADE;")
	if err != nil {
		return err
	}
	return nil
}
