package connectors

import "temsys"

type EcoHTTPConnector struct {
	http HTTPConnector
	scheduler temsys.ReportScheluder
}

type wakeupPeriod struct {
	start, end time.Time
}

const (
	sensorUptime int = 5
)

func NewEcoConnector(http HTTPConnector) EcoHTTPConnector {
	return EcoHTTPConnector{
		http,
	}
}

func (connector EcoHTTPConnector) ReadDataFor(sensor temsys.Sensor) ([]temsys.Report, error) {
	sensor.UpdateInterval
	return connector.http.ReadDataFor(sensor)
}

func (connector EcoHTTPConnector) calculateNextWakeup(interval int64) wakeupPeriod {

}

func 
