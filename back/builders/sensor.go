package builders

import (
	"temsys"
	"temsys/connectors"
)

type HttpSensorBuilder struct {
	name        string
	conntype    temsys.ConnectionType
	connvalue   string
	interval    int64
	reportTypes []string
	deleted     bool
}

func NewHttpSensorBuilder() temsys.SensorBuilder {
	return HttpSensorBuilder{
		name:        "",
		conntype:    "http",
		connvalue:   "127.0.0.1",
		interval:    60,
		reportTypes: []string{},
		deleted:     false,
	}
}

func (builder HttpSensorBuilder) WithName(name string) temsys.SensorBuilder {
	builder.name = name
	return builder
}

func (builder HttpSensorBuilder) WithConnection(conntype temsys.ConnectionType, connvalue string) temsys.SensorBuilder {
	builder.conntype = conntype
	builder.connvalue = connvalue
	return builder
}

func (builder HttpSensorBuilder) WithUpdateInterval(interval int64) temsys.SensorBuilder {
	builder.interval = interval
	return builder
}

func (builder HttpSensorBuilder) WithSupportedReports(types []string) temsys.SensorBuilder {
	builder.reportTypes = types
	return builder
}

func (builder HttpSensorBuilder) IsDeleted(deleted bool) temsys.SensorBuilder {
	builder.deleted = deleted
	return builder
}

func (builder HttpSensorBuilder) Build() temsys.Sensor {
	return temsys.Sensor{
		Name:             builder.name,
		ConnType:         builder.conntype,
		ConnValue:        builder.connvalue,
		UpdateInterval:   builder.interval,
		Deleted:          builder.deleted,
		SupportedReports: builder.reportTypes,
		Connector:        connectors.HTTPConnector{IP: builder.connvalue},
	}
}
