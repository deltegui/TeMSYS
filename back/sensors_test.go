package temsys_test

import (
	"fmt"
	"temsys"
	"temsys/testu"
	"testing"
	"time"
)

type fakeSensorRepo struct {
	connector       temsys.SensorConnector
	saveReturnErr   error
	saveReturnInput temsys.Sensor
	getAllReturn    []temsys.Sensor
	getAllErr       error
	getAllInput     temsys.ShowDeleted
	getByNameInput  string
	getByNameReturn temsys.Sensor
	getByNameErr    error
	updateInput     temsys.Sensor
	updateReturn    bool
}

func (repo *fakeSensorRepo) Save(sensor temsys.Sensor) error {
	repo.saveReturnInput = sensor
	return repo.saveReturnErr
}

func (repo *fakeSensorRepo) GetAll(showDeleted temsys.ShowDeleted) ([]temsys.Sensor, error) {
	repo.getAllInput = showDeleted
	return repo.getAllReturn, repo.getAllErr
}

func (repo *fakeSensorRepo) GetByName(name string) (temsys.Sensor, error) {
	repo.getByNameInput = name
	return repo.getByNameReturn, repo.getByNameErr
}

func (repo *fakeSensorRepo) Update(sensor temsys.Sensor) bool {
	repo.updateInput = sensor
	return repo.updateReturn
}

type fakeSensorConnector struct {
	reports []temsys.Report
	err     error
}

func (conn fakeSensorConnector) ReadDataFor(s temsys.Sensor) ([]temsys.Report, error) {
	return conn.reports, conn.err
}

func TestSensorNowCase(t *testing.T) {
	sensor := temsys.Sensor{
		Name:           "salon",
		ConnType:       temsys.HTTP,
		ConnValue:      "192.168.1.3",
		UpdateInterval: 60,
		Deleted:        false,
		SupportedReports: []string{
			"temperature",
			"humidiy",
		},
	}
	nowReports := []temsys.Report{
		{
			ReportType: "temperature",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:12:01Z"),
			Value:      30.4,
		},
		{
			ReportType: "humidity",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:12:01Z"),
			Value:      38.3,
		},
	}

	t.Run("Should return current sensor status", func(t *testing.T) {
		expected := []temsys.ReportResponse{
			{
				ReportType: "temperature",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      30.4,
			},
			{
				ReportType: "humidity",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      38.3,
			},
		}
		sensor.Connector = fakeSensorConnector{
			reports: nowReports,
			err:     nil,
		}
		sensorRepo := fakeSensorRepo{
			getByNameReturn: sensor,
			getByNameErr:    nil,
		}
		nowCase := temsys.NewSensorNowCase(&sensorRepo)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, "salon")

		res := presenter.Data.([]temsys.ReportResponse)
		testu.Equals(t, sensorRepo.getByNameInput, "salon")
		testu.Equals(t, expected, res)
	})

	t.Run("Should fail if sensor not respond", func(t *testing.T) {
		sensor.Connector = fakeSensorConnector{
			err: fmt.Errorf("Fail to read"),
		}
		sensorRepo := fakeSensorRepo{
			getByNameReturn: sensor,
			getByNameErr:    nil,
		}

		nowCase := temsys.NewSensorNowCase(&sensorRepo)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, "salon")

		res := presenter.DataErr.(error)
		testu.Equals(t, sensorRepo.getByNameInput, "salon")
		testu.Equals(t, temsys.SensorNotRespondErr, res)
	})

	t.Run("Should fail if sensor does not exists", func(t *testing.T) {
		sensor.Connector = fakeSensorConnector{
			err: fmt.Errorf("Fail to read"),
		}
		sensorRepo := fakeSensorRepo{
			getByNameErr: fmt.Errorf("Sensor not found"),
		}

		nowCase := temsys.NewSensorNowCase(&sensorRepo)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, "salon")

		res := presenter.DataErr.(error)
		testu.Equals(t, sensorRepo.getByNameInput, "salon")
		testu.Equals(t, temsys.SensorNotFoundErr, res)
	})
}

type fakeReportCache struct {
	updateInput []temsys.Report
	updateErr   error
	getInput    string
	getReturn   []temsys.Report
	getErr      error
}

func (cache *fakeReportCache) Update(report []temsys.Report) error {
	cache.updateInput = report
	return cache.updateErr
}

func (cache *fakeReportCache) GetForSensor(name string) ([]temsys.Report, error) {
	cache.getInput = name
	return cache.getReturn, cache.getErr
}

func TestCachedNowCase(t *testing.T) {
	sensor := temsys.Sensor{
		Name:           "salon",
		ConnType:       temsys.HTTP,
		ConnValue:      "192.168.1.3",
		UpdateInterval: 60,
		Deleted:        false,
		SupportedReports: []string{
			"temperature",
			"humidiy",
		},
	}
	nowReports := []temsys.Report{
		{
			ReportType: "temperature",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:12:01Z"),
			Value:      30.4,
		},
		{
			ReportType: "humidity",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:12:01Z"),
			Value:      38.3,
		},
	}
	cachedReports := []temsys.Report{
		{
			ReportType: "temperature",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:10:01Z"),
			Value:      27.4,
		},
		{
			ReportType: "humidity",
			SensorName: "salon",
			Date:       testu.TimeParseOrPanic("2022-01-03T13:10:01Z"),
			Value:      20.3,
		},
	}
	fiveMinutes := 5 * time.Minute
	inTime := testu.TimeParseOrPanic("2022-01-03T13:13:13Z")
	outTime := testu.TimeParseOrPanic("2022-01-03T13:20:01Z")

	t.Run("Should always make a sensor request if you are admin", func(t *testing.T) {
		expected := []temsys.ReportResponse{
			{
				ReportType: "temperature",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      30.4,
			},
			{
				ReportType: "humidity",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      38.3,
			},
		}
		sensor.Connector = fakeSensorConnector{
			reports: nowReports,
			err:     nil,
		}
		sensorRepo := fakeSensorRepo{
			getByNameReturn: sensor,
			getByNameErr:    nil,
		}
		clock := testu.FakeClock{NowReturn: outTime}
		cache := fakeReportCache{}
		nowCase := temsys.NewCachedSensorNowCase(clock, &sensorRepo, &cache, fiveMinutes)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, temsys.CachedSensorNowRequest{
			UserRole: temsys.AdminRole,
			Sensor:   "salon",
		})

		res := presenter.Data.([]temsys.ReportResponse)
		testu.Equals(t, sensorRepo.getByNameInput, "salon")
		testu.Equals(t, expected, res)
	})

	t.Run("Should return cached report if you are normal user and they arent outdated", func(t *testing.T) {
		expected := []temsys.ReportResponse{
			{
				ReportType: "temperature",
				SensorName: "salon",
				Date:       "2022-01-03T13:10:01.000Z",
				Value:      27.4,
			},
			{
				ReportType: "humidity",
				SensorName: "salon",
				Date:       "2022-01-03T13:10:01.000Z",
				Value:      20.3,
			},
		}
		sensor.Connector = fakeSensorConnector{
			reports: nowReports,
			err:     nil,
		}
		sensorRepo := fakeSensorRepo{
			getByNameReturn: sensor,
			getByNameErr:    nil,
		}
		cache := fakeReportCache{
			getReturn: cachedReports,
			getErr:    nil,
		}
		clock := testu.FakeClock{NowReturn: inTime}

		nowCase := temsys.NewCachedSensorNowCase(clock, &sensorRepo, &cache, fiveMinutes)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, temsys.CachedSensorNowRequest{
			UserRole: temsys.UserRole,
			Sensor:   "salon",
		})

		res := presenter.Data.([]temsys.ReportResponse)
		testu.Equals(t, cache.getInput, "salon")
		testu.Equals(t, expected, res)
	})

	t.Run("Should update cache and return new reports", func(t *testing.T) {
		expected := []temsys.ReportResponse{
			{
				ReportType: "temperature",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      30.4,
			},
			{
				ReportType: "humidity",
				SensorName: "salon",
				Date:       "2022-01-03T13:12:01.000Z",
				Value:      38.3,
			},
		}
		sensor.Connector = fakeSensorConnector{
			reports: nowReports,
			err:     nil,
		}
		sensorRepo := fakeSensorRepo{
			getByNameReturn: sensor,
			getByNameErr:    nil,
		}
		cache := fakeReportCache{
			getReturn: cachedReports,
			getErr:    nil,
		}
		clock := testu.FakeClock{NowReturn: outTime}
		nowCase := temsys.NewCachedSensorNowCase(clock, &sensorRepo, &cache, 5*time.Minute)
		presenter := testu.FakePresenter{}

		nowCase.Exec(&presenter, temsys.CachedSensorNowRequest{
			UserRole: temsys.UserRole,
			Sensor:   "salon",
		})

		res := presenter.Data.([]temsys.ReportResponse)
		testu.Equals(t, "salon", cache.getInput)
		testu.Equals(t, nowReports, cache.updateInput)
		testu.Equals(t, expected, res)
	})
}
