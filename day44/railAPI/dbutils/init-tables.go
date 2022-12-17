package dbutils

import "log"
import "database/sql"


func Initialize(dbDriver *sql.DB) {
	tables := []string{train, station, schedule}

	for _, tablestr := range tables {
		statement, driverError := dbDriver.Prepare(tablestr)
		if driverError != nil {
			log.Println(driverError)
		}
		_, statementError := statement.Exec()
		if statementError != nil {
			log.Println("Table ", tablestr, " already exists!")
		}
	}

	log.Println("All tables created/initialized successfully!")
}