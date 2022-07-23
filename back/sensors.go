package temsys

import (
	"log"
	"time"
)

// ConnectionType that a sensor can have.
type ConnectionType string

// Supported ConnectionTypes
const (
	HTTP ConnectionType = "http"
)

// SensorConnector is an abstraction over
// the way to communicate to real sensor
type SensorConnector interface {
	ReadDataFor(Sensor) ([]Report, error)
}

type SensorConfigurer interface {
	Configure(old, new Sensor)
}

// Sensor is a physical device that reads envorment
// values. It will return those values as reports of
// different types.
type Sensor struct {
	Name             string         `db:"name"`
	ConnType         ConnectionType `db:"conntype"`
	ConnValue        string         `db:"connvalue"`
	UpdateInterval   int64          `db:"update_interval"`
	Deleted          bool           `db:"deleted"`
	SupportedReports []string
	Connector        SensorConnector
}

// GetCurrentState reads data from sensor.
func (sensor Sensor) GetCurrentState() ([]Report, error) {
	for i := 0; i < 2; i++ {
		reports, err := sensor.Connector.ReadDataFor(sensor)
		if err == nil {
			return reports, nil
		}
	}
	return []Report{}, SensorNotRespondErr
}

// ShowDeleted is a flag that tells SensorRepo if you want deleted sensors or not.
type ShowDeleted bool

// ShowDeleted values
const (
	WithDeletedSensors    ShowDeleted = true
	WithoutDeletedSensors ShowDeleted = false
)

// SensorRepo is an abstraction over a store of sensors.
type SensorRepo interface {
	Save(Sensor) error
	GetAll(showDeleted ShowDeleted) ([]Sensor, error)
	GetByName(name string) (Sensor, error)
	Update(sensor Sensor) bool
}

// SensorBuilder builds sensors. This thing exists because sensors are difficult to
// create, due it have a SensorConnector.
type SensorBuilder interface {
	WithName(string) SensorBuilder
	WithConnection(ConnectionType, string) SensorBuilder
	WithUpdateInterval(int64) SensorBuilder
	WithSupportedReports([]string) SensorBuilder
	IsDeleted(bool) SensorBuilder
	Build() Sensor
}

// ConnectionResponse is the information to get from clients and show about sensors connection.
type ConnectionResponse struct {
	ConnType  ConnectionType `validate:"required" json:"type"`
	ConnValue string         `validate:"required" json:"value"`
}

// SensorResponse is the information to get from clients and show about sensors.
type SensorResponse struct {
	Name             string             `validate:"required" json:"name"`
	Connection       ConnectionResponse `validate:"required" json:"connection"`
	UpdateInterval   int64              `validate:"required" json:"updateInterval"`
	Deleted          bool               `json:"deleted"`
	SupportedReports []string           `validate:"required" json:"supportedReports"`
	SensorBuilder    SensorBuilder      `json:"-"`
}

// ToSensor translates this viewmodel to Sensor domain object.
func (model SensorResponse) ToSensor() Sensor {
	return model.SensorBuilder.
		WithName(model.Name).
		WithConnection(model.Connection.ConnType, model.Connection.ConnValue).
		WithSupportedReports(model.SupportedReports).
		WithUpdateInterval(model.UpdateInterval).
		IsDeleted(model.Deleted).
		Build()
}

// CreateResponseFromSensor transates sensor struct to its response equivalent.
func CreateResponseFromSensor(sensor Sensor) SensorResponse {
	return SensorResponse{
		Name: sensor.Name,
		Connection: ConnectionResponse{
			ConnType:  sensor.ConnType,
			ConnValue: sensor.ConnValue,
		},
		Deleted:          sensor.Deleted,
		UpdateInterval:   sensor.UpdateInterval,
		SupportedReports: sensor.SupportedReports,
	}
}

// ReportResponse is the information to get from clients and show about sensors.
type ReportResponse struct {
	ReportType string  `json:"type"`
	SensorName string  `json:"sensor"`
	Date       string  `json:"date"`
	Value      float32 `json:"value"`
}

func createReportResponse(report Report) ReportResponse {
	date := report.Date.Format("2006-01-02T15:04:05.000Z")
	return ReportResponse{
		ReportType: report.ReportType,
		SensorName: report.SensorName,
		Date:       date,
		Value:      report.Value,
	}
}

func transformReportsToResponse(reports []Report) []ReportResponse {
	var output []ReportResponse
	for _, r := range reports {
		output = append(output, createReportResponse(r))
	}
	return output
}

// GetAllSensorsRequest for GetAllSensors use case
type GetAllSensorsRequest struct {
	WantDeleted bool
}

// GetAllSensorsCase get all sensors in the system.
type GetAllSensorsCase struct {
	sensorRepo SensorRepo
}

// NewGetAllSensorsCase Creates GetAllSensors use case.
func NewGetAllSensorsCase(sensorRepo SensorRepo) UseCase {
	return GetAllSensorsCase{sensorRepo}
}

// Exec GetAllSensors use case.
func (useCase GetAllSensorsCase) Exec(presenter Presenter, req UseCaseRequest) {
	getAllReq := req.(GetAllSensorsRequest)
	sensors, err := useCase.sensorRepo.GetAll(ShowDeleted(getAllReq.WantDeleted))
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	viewModels := []SensorResponse{}
	for _, sensor := range sensors {
		viewModels = append(viewModels, CreateResponseFromSensor(sensor))
	}
	presenter.Present(viewModels)
}

// SaveSensorCase saves a sensor in a repository
type SaveSensorCase struct {
	sensorRepo     SensorRepo
	reportTypeRepo ReportTypeRepo
	validator      Validator
	reporter       Reporter
}

// NewSaveSensorCase creates a SaveSensorCase
func NewSaveSensorCase(validator Validator, sensorRepo SensorRepo, reporter Reporter, reportTypeRepo ReportTypeRepo) UseCase {
	return Validate(SaveSensorCase{
		sensorRepo,
		reportTypeRepo,
		validator,
		reporter,
	}, validator)
}

func haveReportTypes(reportTypeRepo ReportTypeRepo, reportTypes []string) bool {
	foundRTypes := reportTypeRepo.GetAll()
	for _, rType := range reportTypes {
		found := false
		for _, haveRType := range foundRTypes {
			if rType == haveRType {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Exec save sensor use case
func (useCase SaveSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	viewModel := req.(SensorResponse)
	if !haveReportTypes(useCase.reportTypeRepo, viewModel.SupportedReports) {
		presenter.PresentError(ReportTypeDoesNotExists)
		return
	}
	if err := useCase.sensorRepo.Save(viewModel.ToSensor()); err != nil {
		log.Println(err)
		presenter.PresentError(SensorAlreadyExist)
		return
	}
	presenter.Present(viewModel)
	useCase.reporter.Restart()
}

// DeleteSensorCase from a repository.
type DeleteSensorCase struct {
	sensorRepo SensorRepo
	reporter   Reporter
}

// NewDeleteSensorCase creates DeleteSensorCase with a SensorRepo and a reporter.
func NewDeleteSensorCase(sensorRepo SensorRepo, reporter Reporter) UseCase {
	return DeleteSensorCase{
		sensorRepo,
		reporter,
	}
}

// Exec delete sensor case.
func (useCase DeleteSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensorName := req.(string)
	sensor, err := useCase.sensorRepo.GetByName(sensorName)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	sensor.Deleted = true
	if !useCase.sensorRepo.Update(sensor) {
		presenter.PresentError(InternalErr)
		return
	}
	presenter.Present(struct{ Deleted bool }{true})
	useCase.reporter.Restart()
}

// UpdateSensorCase updates an existing sensor with new data.
type UpdateSensorCase struct {
	sensorRepo     SensorRepo
	reporter       Reporter
	reportTypeRepo ReportTypeRepo
}

// NewUpdateSensorCase creates a update sensor case.
func NewUpdateSensorCase(validator Validator, sensorRepo SensorRepo, reporter Reporter, reportTypeRepo ReportTypeRepo) UseCase {
	return Validate(UpdateSensorCase{
		sensorRepo,
		reporter,
		reportTypeRepo,
	}, validator)
}

// Exec update sensor case.
func (useCase UpdateSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	viewModel := req.(SensorResponse)
	if _, err := useCase.sensorRepo.GetByName(viewModel.Name); err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	if !haveReportTypes(useCase.reportTypeRepo, viewModel.SupportedReports) {
		presenter.PresentError(ReportTypeDoesNotExists)
		return
	}
	if !useCase.sensorRepo.Update(viewModel.ToSensor()) {
		presenter.PresentError(UpdateErr)
		return
	}
	presenter.Present(viewModel)
	useCase.reporter.Restart()
}

// SensorNowCase reads information from sensor.
type SensorNowCase struct {
	sensorRepo SensorRepo
}

// NewSensorNowCase creates a SensorNowCase.
func NewSensorNowCase(sensorRepo SensorRepo) UseCase {
	return SensorNowCase{
		sensorRepo,
	}
}

// Exec sensor now case.
func (useCase SensorNowCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensorName := req.(string)
	sensor, err := useCase.sensorRepo.GetByName(sensorName)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	reports, err := sensor.GetCurrentState()
	if err != nil {
		presenter.PresentError(err)
		return
	}
	presenter.Present(transformReportsToResponse(reports))
}

// AllSensorNowCase reads information all sensors in the system.
type AllSensorNowCase struct {
	sensorRepo SensorRepo
}

// NewAllSensorNowCase creates all sensor now case.
func NewAllSensorNowCase(sensorRepo SensorRepo) UseCase {
	return AllSensorNowCase{
		sensorRepo,
	}
}

// Exec all sensor now case.
func (useCase AllSensorNowCase) Exec(presenter Presenter, req UseCaseRequest) {
	reports, err := getAllSensorState(useCase.sensorRepo)
	if err != nil {
		presenter.PresentError(err)
		return
	}
	presenter.Present(transformReportsToResponse(reports))
}

// AllSensorNowAverageCase reads current information from all sensors and
// calculates the average of each report type.
type AllSensorNowAverageCase struct {
	sensorRepo SensorRepo
}

// NewAllSensorNowAverageCase creates all sensor now average.
func NewAllSensorNowAverageCase(sensorRepo SensorRepo) UseCase {
	return AllSensorNowAverageCase{
		sensorRepo,
	}
}

// Exec all sensor now average case.
func (useCase AllSensorNowAverageCase) Exec(presenter Presenter, req UseCaseRequest) {
	reports, err := getAllSensorState(useCase.sensorRepo)
	if err != nil {
		presenter.PresentError(err)
		return
	}
	averages := calculateAverages(reports)
	presenter.Present(transformReportsToResponse(averages))
}

type sumReport struct {
	sum   float32
	count int
}

func calculateAverages(reports []Report) []Report {
	sum := sumReportsByType(reports)
	return createReportsWithAverages(sum)
}

func sumReportsByType(reports []Report) map[string]*sumReport {
	reportTypes := getReportTypesFromReports(reports)
	sum := createSumReportMap(reportTypes)
	for _, report := range reports {
		r := sum[report.ReportType]
		r.sum += report.Value
		r.count++
	}
	return sum
}

func createSumReportMap(rTypes []string) map[string]*sumReport {
	sum := make(map[string]*sumReport)
	for _, tp := range rTypes {
		sum[tp] = &sumReport{
			sum:   0,
			count: 0,
		}
	}
	return sum
}

func createReportsWithAverages(sum map[string]*sumReport) []Report {
	var finalReports []Report
	for key, value := range sum {
		average := float64(value.sum) / float64(value.count)
		finalReports = append(finalReports, Report{
			ReportType: key,
			SensorName: "average",
			Date:       time.Now(),
			Value:      roundReportValue(average),
		})
	}
	return finalReports
}

func getReportTypesFromReports(rr []Report) []string {
	var foundTypes []string
	for _, r := range rr {
		haveIt := false
		for _, t := range foundTypes {
			if t == r.ReportType {
				haveIt = true
				break
			}
		}
		if !haveIt {
			foundTypes = append(foundTypes, r.ReportType)
		}
	}
	return foundTypes
}

func getAllSensorState(sensorRepo SensorRepo) ([]Report, error) {
	sensors, err := sensorRepo.GetAll(WithoutDeletedSensors)
	if err != nil {
		return nil, SensorNotFoundErr
	}
	var reports []Report
	for _, sensor := range sensors {
		r, err := sensor.GetCurrentState()
		if err == nil {
			reports = append(reports, r...)
		}
	}
	return reports, nil
}

// GetSensorCase gets one sensor an presents it.
type GetSensorCase struct {
	sensorRepo SensorRepo
}

// NewGetSensorCase creates GetSensorCase.
func NewGetSensorCase(sensorRepo SensorRepo) UseCase {
	return GetSensorCase{
		sensorRepo,
	}
}

// Exec get sensor case.
func (useCase GetSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensorName := req.(string)
	sensor, err := useCase.sensorRepo.GetByName(sensorName)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	presenter.Present(CreateResponseFromSensor(sensor))
}
