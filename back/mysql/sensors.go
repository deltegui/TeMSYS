package mysql

import (
	"fmt"
	"log"
	"temsys"
	"temsys/connectors"

	"github.com/jmoiron/sqlx"
)

// SqlxSensorRepo implementation for SensorRepo using sqlx and mysql
type SqlxSensorRepo struct {
	sqlRepository
}

// NewSqlxSensorRepo using an existing sqlx.DB
func NewSensorRepo(db *sqlx.DB) temsys.SensorRepo {
	return SqlxSensorRepo{sqlRepository{db}}
}

// Save a sensor. If it fails to save returns an error.
func (repo SqlxSensorRepo) Save(sensor temsys.Sensor) error {
	insertSensor := "insert into sensors (name, conntype, connvalue, update_interval)values($1, $2, $3 ,$4)"
	tx := repo.beginOrFatal()
	if _, err := tx.Exec(insertSensor, sensor.Name, sensor.ConnType, sensor.ConnValue, sensor.UpdateInterval); err != nil {
		tx.Rollback()
		return err
	}
	if err := repo.saveSupportedReportsForSensor(tx, sensor); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Update a sensor. Returns ture if everithing was ok. Return false if it cant be saved.
func (repo SqlxSensorRepo) Update(sensor temsys.Sensor) bool {
	update := "update sensors set conntype = $1, connvalue = $2, update_interval = $3, deleted = $4 where name = $5"
	tx := repo.beginOrFatal()
	defer tx.Commit()
	if _, err := tx.Exec(update, sensor.ConnType, sensor.ConnValue, sensor.UpdateInterval, sensor.Deleted, sensor.Name); err != nil {
		log.Fatal(err)
	}
	if err := repo.saveSupportedReportsForSensor(tx, sensor); err != nil {
		log.Fatal(err)
	}
	return true
}

func (repo SqlxSensorRepo) saveSupportedReportsForSensor(tx *sqlx.Tx, sensor temsys.Sensor) error {
	insertReport := "insert into used_report_types (sensor, report_type, add_date)values($1, $2, now())"
	if _, err := tx.Exec("delete from used_report_types where sensor = $1", sensor.Name); err != nil {
		return err
	}
	for _, reportType := range sensor.SupportedReports {
		if _, err := tx.Exec(insertReport, sensor.Name, reportType); err != nil {
			return err
		}
	}
	return nil
}

// GetAll sensors available in the system.
func (repo SqlxSensorRepo) GetAll(showDeleted temsys.ShowDeleted) ([]temsys.Sensor, error) {
	var sensors []temsys.Sensor
	err := repo.db.Select(&sensors, "select name, conntype, connvalue, update_interval, deleted from sensors where deleted = false or deleted = $1", showDeleted)
	if err != nil {
		log.Println(err)
		return []temsys.Sensor{}, err
	}
	if sensors == nil {
		return []temsys.Sensor{}, nil
	}
	for i := 0; i < len(sensors); i++ {
		repo.fillSupportedReportsForSensor(&sensors[i])
		sensors[i].Connector = connectors.HTTPConnector{IP: sensors[i].ConnValue}
	}
	return sensors, nil
}

// GetByName one sensor. Returns an error if it doesnt exists or there is an error
// fetching it.
func (repo SqlxSensorRepo) GetByName(name string) (temsys.Sensor, error) {
	var sensor []temsys.Sensor
	err := repo.db.Select(&sensor, "select name, conntype, connvalue, update_interval, deleted from sensors where name = $1", name)
	if err != nil || len(sensor) < 1 {
		log.Printf("Sensor not found or error: %s\n", err)
		return temsys.Sensor{}, fmt.Errorf("sensor not found")
	}
	repo.fillSupportedReportsForSensor(&sensor[0])
	sensor[0].Connector = connectors.HTTPConnector{IP: sensor[0].ConnValue}
	return sensor[0], nil
}

func (repo SqlxSensorRepo) fillSupportedReportsForSensor(sensor *temsys.Sensor) {
	var reports []string = []string{}
	err := repo.db.Select(&reports, "select report_type from used_report_types where sensor = $1", sensor.Name)
	if err != nil {
		log.Println(err)
		sensor.SupportedReports = []string{}
		return
	}
	sensor.SupportedReports = reports
}
