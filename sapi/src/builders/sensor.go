package builders

import (
	"sensorapi/src/connectors"
	"sensorapi/src/domain"
)

type HttpSensorBuilder struct {
	name        string
	conntype    domain.ConnectionType
	connvalue   string
	interval    int64
	reportTypes []domain.ReportType
	deleted     bool
}

func NewHttpSensorBuilder() domain.SensorBuilder {
	return HttpSensorBuilder{
		name:        "",
		conntype:    "http",
		connvalue:   "127.0.0.1",
		interval:    60,
		reportTypes: []domain.ReportType{},
		deleted:     false,
	}
}

func (builder HttpSensorBuilder) WithName(name string) domain.SensorBuilder {
	builder.name = name
	return builder
}

func (builder HttpSensorBuilder) WithConnection(conntype domain.ConnectionType, connvalue string) domain.SensorBuilder {
	builder.conntype = conntype
	builder.connvalue = connvalue
	return builder
}

func (builder HttpSensorBuilder) WithUpdateInterval(interval int64) domain.SensorBuilder {
	builder.interval = interval
	return builder
}

func (builder HttpSensorBuilder) WithSupportedReports(types []domain.ReportType) domain.SensorBuilder {
	builder.reportTypes = types
	return builder
}

func (builder HttpSensorBuilder) IsDeleted(deleted bool) domain.SensorBuilder {
	builder.deleted = deleted
	return builder
}

func (builder HttpSensorBuilder) Build() domain.Sensor {
	return domain.Sensor{
		Name:             builder.name,
		ConnType:         builder.conntype,
		ConnValue:        builder.connvalue,
		UpdateInterval:   builder.interval,
		Deleted:          builder.deleted,
		SupportedReports: builder.reportTypes,
		Connector:        connectors.HTTPConnector{IP: builder.connvalue},
	}
}
