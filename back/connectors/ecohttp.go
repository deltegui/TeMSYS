package connectors

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"temsys"
)

type EcoHTTPConnector struct {
	scheduler  temsys.ReportScheluder
	reportRepo temsys.ReportRepo
}

func NewEcoConnector(scheduler temsys.ReportScheluder, repo temsys.ReportRepo) EcoHTTPConnector {
	return EcoHTTPConnector{
		scheduler:  scheduler,
		reportRepo: repo,
	}
}

func (connector EcoHTTPConnector) ReadDataFor(sensor temsys.Sensor) ([]temsys.Report, error) {
	return connector.reportRepo.GetFiltered(temsys.ReportFilter{
		SensorName: sensor.Name,
		Trim:       len(sensor.SupportedReports),
	}), nil
}

func (connect EcoHTTPConnector) Configure(old, new temsys.Sensor) {
	content := fmt.Sprintf(
		"%s;%d;",
		new.ConnValue,
		new.UpdateInterval)
	for supported := range new.SupportedReports {
		content = fmt.Sprintf("%s,%d", content, supported)
	}
	configureSensor := func() {
		_, err := http.Post(old.ConnValue, "application/sensor_config", strings.NewReader(content))
		if err != nil {
			log.Printf(
				"Error sending configuration to sensor '%s' with IP '%s': %s",
				old.ConnValue,
				old.Name,
				err)
			return
		}
		log.Printf(
			"Changed sensor configuration '%s' with IP '%s'",
			old.ConnValue,
			old.Name)
	}
	connect.scheduler.AddJobOnce(configureSensor, old.UpdateInterval)
}
