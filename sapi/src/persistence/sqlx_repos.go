package persistence

import (
	"fmt"
	"log"
	"sensorapi/src/connectors"
	"sensorapi/src/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type SqlxReportTypeRepo struct {
	db *sqlx.DB
}

func NewSqlxReportTypeRepo(conn *SqlxConnection) domain.ReportTypeRepo {
	return SqlxReportTypeRepo{conn.GetConnection()}
}

func (repo SqlxReportTypeRepo) Save(rType domain.ReportType) error {
	insert := "insert into REPORT_TYPES values(?)"
	tx, err := repo.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec(insert, rType)
	return err
}

func (repo SqlxReportTypeRepo) GetAll() []domain.ReportType {
	tx, err := repo.db.Beginx()
	if err != nil {
		return []domain.ReportType{}
	}
	defer tx.Commit()
	var types []domain.ReportType
	err = tx.Select(&types, "SELECT TYPE_NAME FROM REPORT_TYPES")
	if err != nil {
		log.Println(err)
		return []domain.ReportType{}
	}
	return types
}

type SqlxSensorRepo struct {
	db *sqlx.DB
}

func NewSqlxSensorRepo(conn *SqlxConnection) domain.SensorRepo {
	return SqlxSensorRepo{conn.GetConnection()}
}

func (repo SqlxSensorRepo) Save(sensor domain.Sensor) error {
	insertSensor := "insert into SENSORS (NAME, CONNTYPE, CONNVALUE, UPDATE_INTERVAL)values(?, ?, ? ,?)"
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	if _, err = tx.Exec(insertSensor, sensor.Name, sensor.ConnType, sensor.ConnValue, sensor.UpdateInterval); err != nil {
		return err
	}
	return repo.saveSupportedReportsForSensor(tx, sensor)
}

func (repo SqlxSensorRepo) saveSupportedReportsForSensor(tx *sqlx.Tx, sensor domain.Sensor) error {
	insertReport := "insert into USED_REPORT_TYPES (SENSOR, REPORT_TYPE, ADD_DATE)values(?, ?, NOW())"
	if _, err := tx.Exec("DELETE FROM USED_REPORT_TYPES WHERE SENSOR LIKE ?", sensor.Name); err != nil {
		return err
	}
	for _, reportType := range sensor.SupportedReports {
		if _, err := tx.Exec(insertReport, sensor.Name, reportType); err != nil {
			return err
		}
	}
	return nil
}

func (repo SqlxSensorRepo) GetAll(showDeleted domain.ShowDeleted) ([]domain.Sensor, error) {
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	var sensors []domain.Sensor
	err = tx.Select(&sensors, "SELECT NAME, CONNTYPE, CONNVALUE, UPDATE_INTERVAL, DELETED FROM SENSORS WHERE DELETED = ?", showDeleted)
	if err != nil {
		log.Println(err)
		return []domain.Sensor{}, err
	}
	if sensors == nil {
		return []domain.Sensor{}, nil
	}
	for i := 0; i < len(sensors); i++ {
		repo.FillSupportedReportsForSensor(tx, &sensors[i])
		sensors[i].Connector = connectors.HTTPConnector{IP: sensors[i].ConnValue}
	}
	return sensors, nil
}

func (repo SqlxSensorRepo) FillSupportedReportsForSensor(tx *sqlx.Tx, sensor *domain.Sensor) {
	var reports []domain.ReportType = []domain.ReportType{}
	err := tx.Select(&reports, "SELECT REPORT_TYPE FROM USED_REPORT_TYPES WHERE SENSOR LIKE ?", sensor.Name)
	if err != nil {
		log.Println(err)
		sensor.SupportedReports = []domain.ReportType{}
		return
	}
	sensor.SupportedReports = reports
}

func (repo SqlxSensorRepo) GetByName(name string) (domain.Sensor, error) {
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	var sensor []domain.Sensor
	err = tx.Select(&sensor, "SELECT NAME, CONNTYPE, CONNVALUE, UPDATE_INTERVAL, DELETED FROM SENSORS WHERE NAME LIKE ?", name)
	if err != nil || len(sensor) < 1 {
		log.Printf("Sensor not found or error: %s\n", err)
		return domain.Sensor{}, fmt.Errorf("Sensor not found")
	}
	repo.FillSupportedReportsForSensor(tx, &sensor[0])
	sensor[0].Connector = connectors.HTTPConnector{IP: sensor[0].ConnValue}
	return sensor[0], nil
}

func (repo SqlxSensorRepo) Update(sensor domain.Sensor) bool {
	update := "UPDATE SENSORS SET CONNTYPE = ?, CONNVALUE = ?, UPDATE_INTERVAL = ?, DELETED = ? WHERE NAME LIKE ?"
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	if _, err = tx.Exec(update, sensor.ConnType, sensor.ConnValue, sensor.UpdateInterval, sensor.Deleted, sensor.Name); err != nil {
		log.Fatal(err)
	}
	if err = repo.saveSupportedReportsForSensor(tx, sensor); err != nil {
		log.Fatal(err)
	}
	return true
}

type SqlxReportRepo struct {
	db *sqlx.DB
}

func NewSqlxReportRepo(conn *SqlxConnection) domain.ReportRepo {
	return SqlxReportRepo{conn.GetConnection()}
}

func (repo SqlxReportRepo) Save(report domain.Report) {
	insert := "INSERT INTO REPORTS (SENSOR, TYPE, VALUE, REPORT_DATE) VALUES(?, ?, ?, ?)"
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	_, err = tx.Exec(insert, report.SensorName, report.ReportType, report.Value, report.Date.UTC())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo SqlxReportRepo) GetAll() []domain.Report {
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	var reports []domain.Report
	err = tx.Select(&reports, "SELECT R.SENSOR, R.TYPE, R.VALUE, R.REPORT_DATE FROM REPORTS AS R, SENSORS AS S WHERE S.NAME LIKE R.SENSOR AND S.DELETED = 0")
	if err != nil {
		log.Fatal(err)
	}
	if reports == nil {
		return []domain.Report{}
	}
	return reports
}

func (repo SqlxReportRepo) GetBetweenDates(from time.Time, to time.Time) []domain.Report {
	dateFormat := "2006-01-02 15:04:05"
	tx, err := repo.db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	var reports []domain.Report
	err = tx.Select(&reports, "SELECT R.SENSOR, R.TYPE, R.VALUE, R.REPORT_DATE FROM REPORTS AS R, SENSORS AS S WHERE REPORT_DATE > ? AND REPORT_DATE < ? AND S.NAME LIKE R.SENSOR AND S.DELETED = 0", from.UTC().Format(dateFormat), to.UTC().Format(dateFormat))
	if err != nil {
		log.Fatal(err)
	}
	if reports == nil {
		return []domain.Report{}
	}
	return reports
}
