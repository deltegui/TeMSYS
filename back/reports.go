package temsys

import (
	"log"
	"time"
)

// Known values of report types
const (
	TemperatureReportType string = "temperature"
	HumidityReportType    string = "humidity"
)

// ReportTypeRepo is a place to store ReportTypes.
// ReportTypes are not hard-coded because we want clients to have the option
// to change available report types at runtime.
type ReportTypeRepo interface {
	Save(string) error
	GetAll() []string
}

// Report represents a Sensor report, generated when you read it's state.
type Report struct {
	ReportType string    `db:"type" json:"type"`
	SensorName string    `db:"sensor" json:"sensor"`
	Date       time.Time `db:"report_date" json:"date"`
	Value      float32   `db:"value" json:"value"`
}

// ReportFilter stores all information to filter reports using a ReportRepo.
type ReportFilter struct {
	From       time.Time
	To         time.Time
	Trim       int
	Type       string
	SensorName string
}

// ReportRepo is a place to store reports.
type ReportRepo interface {
	Save(Report)
	GetAll() []Report
	GetFiltered(filter ReportFilter) []Report
}

// ScheluderJob is a job to be done by the ReportScheluder.
type ScheluderJob func()

// ReportScheluder is an abstraction of something that runs a job every
// amount of time. The start method must be asynchronous, so must start
// a goroutine.
type ReportScheluder interface {
	AddJobEvery(ScheluderJob, int64)
	Start()
	Stop()
}

// FilteredReportsRequest is the request for GetReportByDatesCase.
// Stores a date range of the reports that you want
type FilteredReportsRequest struct {
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	Average    bool      `json:"average"`
	Trim       int       `json:"trim"`
	Type       string    `json:"type"`
	SensorName string    `json:"sensor"`
}

// GetFilteredReports get all reports generated between two dates.
type GetFilteredReports struct {
	reportRepo ReportRepo
}

// NewGetFilteredReports creates new GetReportsByDates use case with validation
func NewGetFilteredReports(validator Validator, reportRepo ReportRepo) UseCase {
	return Validate(
		GetFilteredReports{reportRepo},
		validator,
	)
}

// Exec GetFilteredReports
func (useCase GetFilteredReports) Exec(presenter Presenter, req UseCaseRequest) {
	datesReq := req.(FilteredReportsRequest)
	result := useCase.reportRepo.GetFiltered(ReportFilter{
		From:       datesReq.From,
		To:         datesReq.To,
		Trim:       datesReq.Trim,
		Type:       datesReq.Type,
		SensorName: datesReq.SensorName,
	})
	if datesReq.Average {
		result = calculateAverages(result)
	}
	presenter.Present(result)
}

func newReportReaderSchedulerJob(sensor Sensor, reportRepo ReportRepo) ScheluderJob {
	return func() {
		log.Printf("Running job for %s\n", sensor.Name)
		currentReports, err := sensor.GetCurrentState()
		if err != nil {
			log.Println(err)
			return
		}
		for _, report := range currentReports {
			reportRepo.Save(report)
		}
	}
}

// TODO move this shit to cmd temsys binary
var globalReporter *Reporter

// Reporter creates Reports every time interval using a ReportScheluder.
type Reporter struct {
	sensorRepo SensorRepo
	reportRepo ReportRepo
	scheluder  ReportScheluder
	restart    chan bool
}

// NewReporter creates a reporter
func NewReporter(sensorRepo SensorRepo, reportRepo ReportRepo, scheluder ReportScheluder) Reporter {
	if globalReporter == nil {
		log.Println("Created reporter")
		globalReporter = &Reporter{
			sensorRepo: sensorRepo,
			reportRepo: reportRepo,
			scheluder:  scheluder,
			restart:    make(chan bool),
		}
	}
	return *globalReporter
}

// Start scheduler to retrieve reports.
func (reporter Reporter) Start() {
	for {
		sensors, err := reporter.sensorRepo.GetAll(WithoutDeletedSensors)
		if err != nil {
			return
		}
		for _, sensor := range sensors {
			job := newReportReaderSchedulerJob(sensor, reporter.reportRepo)
			reporter.scheluder.AddJobEvery(job, sensor.UpdateInterval)
		}
		reporter.scheluder.Start()
		<-reporter.restart
		log.Println("Restarting scheduler...")
		reporter.scheluder.Stop()
		log.Println("DONE! Scheduler restarted!")
	}
}

// Restart scheduler.
func (reporter Reporter) Restart() {
	reporter.restart <- true
}
