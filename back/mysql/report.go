package mysql

import (
	"log"
	"temsys"

	"github.com/jmoiron/sqlx"
)

// SqlxReportTypeRepo impelmentation of ReportTypeRepo using sqlx and mysql.
// ReportTypes are not hard-coded because we want clients to have the option
// to change available report types at runtime.
type SqlxReportTypeRepo struct {
	sqlRepository
}

// NewReportTypeRepo with sqlx
func NewReportTypeRepo(db *sqlx.DB) SqlxReportTypeRepo {
	return SqlxReportTypeRepo{sqlRepository{db}}
}

// Save a ReportType. Rarely returns an error.
func (repo SqlxReportTypeRepo) Save(rType string) error {
	insert := "insert into report_types values(?)"
	_, err := repo.db.Exec(insert, rType)
	return err
}

// GetAll report types currently available.
func (repo SqlxReportTypeRepo) GetAll() []string {
	var types []string
	err := repo.db.Select(&types, "select type_name from report_types")
	if err != nil {
		log.Println(err)
		return []string{}
	}
	if types == nil {
		return []string{}
	}
	return types
}

// SqlxReportRepo impelmentation of ReportTypeRepo using sqlx and mysql.
type SqlxReportRepo struct {
	sqlRepository
}

// NewReportRepo creates a new ReportRepo using sqlx.
func NewReportRepo(db *sqlx.DB) temsys.ReportRepo {
	return SqlxReportRepo{sqlRepository{db}}
}

// Save a report.
func (repo SqlxReportRepo) Save(report temsys.Report) {
	insert := "insert into reports (sensor, type, value, report_date) values(?, ?, ?, ?)"
	_, err := repo.db.Exec(insert, report.SensorName, report.ReportType, report.Value, report.Date.UTC())
	if err != nil {
		log.Fatal(err)
	}
}

// GetAll available reports in the system.
func (repo SqlxReportRepo) GetAll() []temsys.Report {
	var reports []temsys.Report
	err := repo.db.Select(&reports, "select r.sensor, r.type, r.value, r.report_date from reports as r, sensors as s where s.name like r.sensor and s.deleted = 0")
	if err != nil {
		log.Fatal(err)
	}
	if reports == nil {
		return []temsys.Report{}
	}
	return reports
}

// GetFiltered let you get all reports using filters.
func (repo SqlxReportRepo) GetFiltered(filter temsys.ReportFilter) []temsys.Report {
	dateFormat := "2006-01-02 15:04:05"
	var reports []temsys.Report
	var err error
	if len(filter.Type) == 0 {
		query := "select r.sensor, r.type, r.value, r.report_date from reports as r, sensors as s where report_date > ? and report_date < ? and s.name like r.sensor and s.deleted = 0 and r.sensor = ? order by r.report_date desc limit ?"
		err = repo.db.Select(&reports, query, filter.From.UTC().Format(dateFormat), filter.To.UTC().Format(dateFormat), filter.SensorName, filter.Trim)
	} else {
		query := "select r.sensor, r.type, r.value, r.report_date from reports as r, sensors as s where report_date > ? and report_date < ? and r.type = ? and s.name like r.sensor and s.deleted = 0 and r.sensor = ? order by r.report_date desc limit ?;"
		err = repo.db.Select(&reports, query, filter.From.UTC().Format(dateFormat), filter.To.UTC().Format(dateFormat), filter.Type, filter.SensorName, filter.Trim)
	}
	if err != nil {
		log.Fatal(err)
	}
	if reports == nil {
		return []temsys.Report{}
	}
	return reports
}

// GetFilteredAverage let you get all reports average using filters.
func (repo SqlxReportRepo) GetFilteredAverage(filter temsys.ReportFilter) []temsys.Report {
	dateFormat := "2006-01-02 15:04:05"
	var reports []temsys.Report
	query := "select 'average' as sensor, r.type, avg(r.value) as value, now() as report_date from reports as r, sensors as s where report_date > ? and report_date < ? and s.name like r.sensor and s.deleted = 0 group by type"
	err := repo.db.Select(&reports, query, filter.From.UTC().Format(dateFormat), filter.To.UTC().Format(dateFormat))
	if err != nil {
		log.Fatal(err)
	}
	if reports == nil {
		return []temsys.Report{}
	}
	return reports
}
