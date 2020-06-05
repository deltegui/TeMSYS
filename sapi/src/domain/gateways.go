package domain

import "time"

type Presenter interface {
	Present(data interface{})
	PresentError(data error)
}

type UseCaseRequest interface{}

var EmptyRequest UseCaseRequest = struct{}{}

type UseCase interface {
	Exec(Presenter, UseCaseRequest)
}

type Validator interface {
	Validate(interface{}) error
}

type ReportType string

type ReportTypeRepo interface {
	Save(ReportType) error
	GetAll() []ReportType
}

type ConnectionType string

const (
	HTTP ConnectionType = "http"
)

type Sensor struct {
	Name             string         `db:"NAME"`
	ConnType         ConnectionType `db:"CONNTYPE"`
	ConnValue        string         `db:"CONNVALUE"`
	UpdateInterval   int64          `db:"UPDATE_INTERVAL"`
	Deleted          bool           `db:"DELETED"`
	SupportedReports []ReportType
	Connector        SensorConnector
}

func (sensor Sensor) GetCurrentState() ([]Report, error) {
	for i := 0; i < 2; i++ {
		reports, err := sensor.Connector.ReadDataFor(sensor)
		if err == nil {
			return reports, nil
		}
	}
	return []Report{}, SensorNotRespondErr
}

type ShowDeleted bool

const (
	WithDeletedSensors    ShowDeleted = true
	WithoutDeletedSensors ShowDeleted = false
)

type SensorRepo interface {
	Save(Sensor) error
	GetAll(showDeleted ShowDeleted) ([]Sensor, error)
	GetByName(name string) (Sensor, error)
	Update(sensor Sensor) bool
}

// Report represents a Sensor report.
type Report struct {
	ReportType string    `db:"TYPE" json:"type"`
	SensorName string    `db:"SENSOR" json:"sensor"`
	Date       time.Time `db:"REPORT_DATE" json:"date"`
	Value      float32   `db:"VALUE" json:"value"`
}

// SensorConnector is an abstraction over
// the way to communicate to real sensor
type SensorConnector interface {
	ReadDataFor(Sensor) ([]Report, error)
}

type ReportRepo interface {
	Save(Report)
	GetAll() []Report
	GetBetweenDates(from time.Time, to time.Time) []Report
}

type ScheluderJob func()

type ReportScheluder interface {
	AddJobEvery(ScheluderJob, int64)
	Start() // DEBE SER NO BLOQUEANTE (QUE ARRANQUE UNA CORRUTINA)
	Stop()
}

type SensorBuilder interface {
	WithName(string) SensorBuilder
	WithConnection(ConnectionType, string) SensorBuilder
	WithUpdateInterval(int64) SensorBuilder
	WithSupportedReports([]ReportType) SensorBuilder
	IsDeleted(bool) SensorBuilder
	Build() Sensor
}

type ReportQueue interface {
	Connect()
	Publish(Report)
}
