package temsys_test

import (
	"temsys"
	"temsys/testu"
	"testing"
	"time"
)

type fakeReportTypeRepo struct {
	saveReturnErr error
	saveInput     string
	getAllReturn  []string
}

func (repo *fakeReportTypeRepo) Save(s string) error {
	repo.saveInput = s
	return repo.saveReturnErr
}

func (repo fakeReportTypeRepo) GetAll() []string {
	return repo.getAllReturn
}

type fakeReportRepo struct {
	saveInput            temsys.Report
	getAllReturn         []temsys.Report
	filteredInput        temsys.ReportFilter
	getFiltered          []temsys.Report
	filteredAverageInput temsys.ReportFilter
	getFilteredAverage   []temsys.Report
}

func (repo *fakeReportRepo) Save(r temsys.Report) {
	repo.saveInput = r
}

func (repo fakeReportRepo) GetAll() []temsys.Report {
	return repo.getAllReturn
}

func (repo *fakeReportRepo) GetFiltered(filter temsys.ReportFilter) []temsys.Report {
	repo.filteredInput = filter
	return repo.getFiltered
}

func (repo *fakeReportRepo) GetFilteredAverage(filter temsys.ReportFilter) []temsys.Report {
	repo.filteredAverageInput = filter
	return repo.getFilteredAverage
}

type fakeReportScheduler struct{}

func (scheduler fakeReportScheduler) AddJobEvery(temsys.ScheluderJob, int64) {}
func (scheduler fakeReportScheduler) Start()                                 {}
func (scheduler fakeReportScheduler) Stop()                                  {}

func reportsIgnoreTime(rr *[]temsys.Report) {
	for i := range *rr {
		r := &(*rr)[i]
		r.Date = time.Unix(0, 0)
	}
}

func TestGetFilteredReportsCase(t *testing.T) {
	expectedFilter := temsys.ReportFilter{
		From:       testu.TimeParseOrPanic("2020-01-02T01:01:01Z"),
		To:         testu.TimeParseOrPanic("2020-01-03T01:01:01Z"),
		Trim:       10,
		Type:       "temperature",
		SensorName: "habitacion",
	}
	caseReq := temsys.FilteredReportsBySensorRequest{
		FilteredReportsRequest: temsys.FilteredReportsRequest{
			From: testu.TimeParseOrPanic("2020-01-02T01:01:01Z"),
			To:   testu.TimeParseOrPanic("2020-01-03T01:01:01Z"),
		},
		Trim:       10,
		Average:    false,
		Type:       "temperature",
		SensorName: "habitacion",
	}

	t.Run("Should call ReportRepo with filters", func(t *testing.T) {
		expectedResult := []temsys.Report{
			{
				ReportType: "temperature",
				SensorName: "habitacion",
				Date:       testu.TimeParseOrPanic("2020-01-02T05:01:01Z"),
				Value:      20,
			},
		}
		presenter := testu.FakePresenter{}
		reportRepo := fakeReportRepo{
			getFiltered: expectedResult,
		}
		filteredCase := temsys.NewGetFilteredReportsBySensor(testu.AlwaysValidValidator{}, &reportRepo)
		filteredCase.Exec(&presenter, caseReq)
		res := presenter.Data.([]temsys.Report)
		testu.Equals(t, expectedFilter, reportRepo.filteredInput)
		testu.Equals(t, res[0], expectedResult[0])
	})

	t.Run("Should calculate report average if is requested", func(t *testing.T) {
		reports := []temsys.Report{
			{
				ReportType: "temperature",
				SensorName: "habitacion",
				Date:       testu.TimeParseOrPanic("2020-01-02T05:01:01Z"),
				Value:      20,
			},
			{
				ReportType: "humidity",
				SensorName: "habitacion",
				Date:       testu.TimeParseOrPanic("2020-01-02T05:01:01Z"),
				Value:      68,
			},
			{
				ReportType: "temperature",
				SensorName: "habitacion",
				Date:       testu.TimeParseOrPanic("2020-01-02T12:01:01Z"),
				Value:      25,
			},
			{
				ReportType: "humidity",
				SensorName: "habitacion",
				Date:       testu.TimeParseOrPanic("2020-01-02T12:01:01Z"),
				Value:      70,
			},
		}
		expectedResult := []temsys.Report{
			{
				ReportType: "temperature",
				SensorName: "average",
				Date:       time.Unix(0, 0),
				Value:      22.5,
			},
			{
				ReportType: "humidity",
				SensorName: "average",
				Date:       time.Unix(0, 0),
				Value:      69,
			},
		}
		presenter := testu.FakePresenter{}
		reportRepo := fakeReportRepo{
			getFiltered: reports,
		}
		filteredCase := temsys.NewGetFilteredReportsBySensor(testu.AlwaysValidValidator{}, &reportRepo)
		averageReq := caseReq
		averageReq.Average = true
		filteredCase.Exec(&presenter, averageReq)
		res := presenter.Data.([]temsys.Report)
		reportsIgnoreTime(&res)
		testu.Equals(t, res, expectedResult)
	})
}

func TestReportIsOlder(t *testing.T) {
	//               Year  Month         D  H   M   S  s
	now := time.Date(2022, time.January, 3, 13, 12, 1, 0, time.Now().Location())
	clock := testu.FakeClock{NowReturn: now}
	report := temsys.Report{
		ReportType: "temperature",
		SensorName: "habitacion",
		Date:       testu.TimeParseOrPanic("2022-01-03T13:01:01Z"),
		Value:      20,
	}
	testu.Equals(t, report.IsOlder(clock, 5*time.Minute), true)
}
