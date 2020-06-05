package domain

import (
	"log"
)

type ConnetionViewModel struct {
	ConnType  ConnectionType `validate:"required" json:"type"`
	ConnValue string         `validate:"required" json:"value"`
}

type SensorViewModel struct {
	Name             string             `validate:"required" json:"name"`
	Connection       ConnetionViewModel `validate:"required" json:"connection"`
	UpdateInterval   int64              `validate:"required" json:"updateInterval"`
	Deleted          bool               `json:"deleted"`
	SupportedReports []ReportType       `validate:"required" json:"supportedReports"`
	SensorBuilder    SensorBuilder      `json:"-"`
}

func (model SensorViewModel) ToSensor() Sensor {
	return model.SensorBuilder.
		WithName(model.Name).
		WithConnection(model.Connection.ConnType, model.Connection.ConnValue).
		WithSupportedReports(model.SupportedReports).
		WithUpdateInterval(model.UpdateInterval).
		IsDeleted(model.Deleted).
		Build()
}

func CreateViewModelFromSensor(sensor Sensor) SensorViewModel {
	return SensorViewModel{
		Name: sensor.Name,
		Connection: ConnetionViewModel{
			ConnType:  sensor.ConnType,
			ConnValue: sensor.ConnValue,
		},
		Deleted:          sensor.Deleted,
		UpdateInterval:   sensor.UpdateInterval,
		SupportedReports: sensor.SupportedReports,
	}
}

type GetAllRequest struct {
	WantDeleted bool
}

type GetAllSensorsCase struct {
	sensorRepo SensorRepo
}

func NewGetAllSensorsCase(sensorRepo SensorRepo) GetAllSensorsCase {
	return GetAllSensorsCase{sensorRepo}
}

func (useCase GetAllSensorsCase) Exec(presenter Presenter, req UseCaseRequest) {
	getAllReq := req.(GetAllRequest)
	var sensors []Sensor
	var err error
	if getAllReq.WantDeleted {
		sensors, err = useCase.sensorRepo.GetAll(WithDeletedSensors)
	} else {
		sensors, err = useCase.sensorRepo.GetAll(WithoutDeletedSensors)
	}
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	viewModels := []SensorViewModel{}
	for _, sensor := range sensors {
		viewModels = append(viewModels, CreateViewModelFromSensor(sensor))
	}
	presenter.Present(viewModels)
}

type SaveSensorCase struct {
	sensorRepo     SensorRepo
	reportTypeRepo ReportTypeRepo
	validator      Validator
	reporter       Reporter
}

func NewSaveSensorCase(sensorRepo SensorRepo, validator Validator, reporter Reporter, reportTypeRepo ReportTypeRepo) SaveSensorCase {
	return SaveSensorCase{
		sensorRepo,
		reportTypeRepo,
		validator,
		reporter,
	}
}

func HaveReportTypes(reportTypeRepo ReportTypeRepo, reportTypes []ReportType) bool {
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

func (useCase SaveSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	viewModel := req.(SensorViewModel)
	if err := useCase.validator.Validate(viewModel); err != nil {
		presenter.PresentError(MalformedRequestErr)
		return
	}
	if !HaveReportTypes(useCase.reportTypeRepo, viewModel.SupportedReports) {
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

type DeleteSensorCase struct {
	sensorRepo SensorRepo
	validator  Validator
	reporter   Reporter
}

func NewDeleteSensorCase(sensorRepo SensorRepo, validator Validator, reporter Reporter) DeleteSensorCase {
	return DeleteSensorCase{
		sensorRepo,
		validator,
		reporter,
	}
}

func (useCase DeleteSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensorName := req.(string)
	sensor, err := useCase.sensorRepo.GetByName(sensorName)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	sensor.Deleted = true
	if useCase.sensorRepo.Update(sensor) {
		presenter.Present(struct{ Deleted bool }{true})
		return
	}
	presenter.PresentError(InternalErr)
	useCase.reporter.Restart()
}

type UpdateSensorCase struct {
	sensorRepo     SensorRepo
	validator      Validator
	reporter       Reporter
	reportTypeRepo ReportTypeRepo
}

func NewUpdateSensorCase(sensorRepo SensorRepo, validator Validator, reporter Reporter, reportTypeRepo ReportTypeRepo) UpdateSensorCase {
	return UpdateSensorCase{
		sensorRepo,
		validator,
		reporter,
		reportTypeRepo,
	}
}

func (useCase UpdateSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	viewModel := req.(SensorViewModel)
	if err := useCase.validator.Validate(viewModel); err != nil {
		log.Println(err)
		presenter.PresentError(MalformedRequestErr)
		return
	}
	if !HaveReportTypes(useCase.reportTypeRepo, viewModel.SupportedReports) {
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

type SensorNowCase struct {
	sensorRepo SensorRepo
	validator  Validator
	reporter   Reporter
}

func NewSensorNowCase(sensorRepo SensorRepo, validator Validator, reporter Reporter) SensorNowCase {
	return SensorNowCase{
		sensorRepo,
		validator,
		reporter,
	}
}

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
	presenter.Present(reports)
}

type AllSensorNowCase struct {
	sensorRepo SensorRepo
}

func NewAllSensorNowCase(sensorRepo SensorRepo) AllSensorNowCase {
	return AllSensorNowCase{
		sensorRepo,
	}
}

func (useCase AllSensorNowCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensors, err := useCase.sensorRepo.GetAll(WithoutDeletedSensors)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	var reports []Report
	for _, sensor := range sensors {
		r, err := sensor.GetCurrentState()
		if err == nil {
			reports = append(reports, r...)
		}
	}
	presenter.Present(reports)
}

type GetSensorCase struct {
	sensorRepo SensorRepo
}

func NewGetSensorCase(sensorRepo SensorRepo) GetSensorCase {
	return GetSensorCase{
		sensorRepo,
	}
}

func (useCase GetSensorCase) Exec(presenter Presenter, req UseCaseRequest) {
	sensorName := req.(string)
	sensor, err := useCase.sensorRepo.GetByName(sensorName)
	if err != nil {
		presenter.PresentError(SensorNotFoundErr)
		return
	}
	presenter.Present(CreateViewModelFromSensor(sensor))
}
