package temsys_test

import (
	"fmt"
	"temsys"
	"temsys/testu"
	"testing"
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

		nowCase.Exec(&presenter, temsys.SensorNowRequest{
			UserRole: temsys.AdminRole,
			Sensor:   "salon",
		})

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

		nowCase.Exec(&presenter, temsys.SensorNowRequest{
			UserRole: temsys.AdminRole,
			Sensor:   "salon",
		})

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

		nowCase.Exec(&presenter, temsys.SensorNowRequest{
			UserRole: temsys.AdminRole,
			Sensor:   "salon",
		})

		res := presenter.DataErr.(error)
		testu.Equals(t, sensorRepo.getByNameInput, "salon")
		testu.Equals(t, temsys.SensorNotFoundErr, res)
	})
}
