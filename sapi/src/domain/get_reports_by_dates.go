package domain

import (
	"time"
)

type ReportsByDatesRequest struct {
	From time.Time `validate:"required" json:"from"`
	To   time.Time `validate:"required" json:"to"`
}

type GetReportsByDates struct {
	reportRepo ReportRepo
	validator  Validator
}

func NewGetReportsByDates(reportRepo ReportRepo, validator Validator) GetReportsByDates {
	return GetReportsByDates{reportRepo, validator}
}

func (useCase GetReportsByDates) Exec(presenter Presenter, req UseCaseRequest) {
	datesReq := req.(ReportsByDatesRequest)
	if err := useCase.validator.Validate(datesReq); err != nil {
		presenter.PresentError(MalformedRequestErr)
		return
	}
	presenter.Present(useCase.reportRepo.GetBetweenDates(datesReq.From, datesReq.To))
}
