package com.deltegui.temsys.sensors.infraestructure;

import com.deltegui.temsys.reports.Report;
import com.deltegui.temsys.sensors.application.SensorRepository;
import com.deltegui.temsys.sensors.domain.Sensor;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.function.Function;
import java.util.stream.Collectors;

@Repository
public class SapiSensorRepository implements SensorRepository {
    private final ObjectMapper mapper;
    private final SensorApi sapiService;

    private final static class Connection {
        private String type;
        private String value;

        public void setType(String type) {
            this.type = type;
        }

        public void setValue(String value) {
            this.value = value;
        }
    }

    private final static class SensorApiModel {
        private String name;
        private Connection connection;
        private int updateInterval;
        private boolean deleted;
        private String[] supportedReports;

        public void setName(String name) {
            this.name = name;
        }

        public void setConnection(Connection connection) {
            this.connection = connection;
        }

        public void setUpdateInterval(int updateInterval) {
            this.updateInterval = updateInterval;
        }

        public void setDeleted(boolean deleted) {
            this.deleted = deleted;
        }

        public void setSupportedReports(String[] supportedReports) {
            this.supportedReports = supportedReports;
        }

        public Sensor toDomain() {
            return new Sensor(this.name, this.connection.type, this.connection.value, this.updateInterval, this.deleted);
        }
    }

    private final static class ReportModel {
        private String type;
        private double value;
        private LocalDateTime date;
        private String sensor;

        public void setType(String type) {
            this.type = type;
        }

        public void setValue(double value) {
            this.value = value;
        }

        public void setDate(LocalDateTime date) {
            this.date = date;
        }

        public void setSensor(String sensor) {
            this.sensor = sensor;
        }

        public Report toDomain() {
            return new Report(this.type, this.value, this.date, this.sensor);
        }
    }

    public SapiSensorRepository(SensorApi sapiService) {
        this.sapiService = sapiService;
        this.mapper = new ObjectMapper().registerModule(new JavaTimeModule());
    }

    public List<Report> getStatusForSensor(String name) {
        return this.getList("/sensor/" + name + "/now",
                res -> Arrays.stream(parseJsonList(res, ReportModel[].class))
                        .map(ReportModel::toDomain)
                        .collect(Collectors.toList()));
    }

    public List<Report> getStatusForAll() {
        return this.getList("/sensors/now",
                res -> Arrays.stream(parseJsonList(res, ReportModel[].class))
                        .map(ReportModel::toDomain)
                        .collect(Collectors.toList()));
    }

    public List<Sensor> getAll() {
        return this.getList("/sensors",
                res -> Arrays.stream(parseJsonList(res, SensorApiModel[].class))
                        .map(SensorApiModel::toDomain)
                        .collect(Collectors.toList()));
    }

    private <T> List<T> getList(String endpoint, Function<String, List<T>> mapReponse) {
        Optional<String> optionalResponse = this.sapiService.get(endpoint);
        if(optionalResponse.isEmpty()) {
            return new ArrayList<>();
        }
        return mapReponse.apply(optionalResponse.get());
    }

    private <T> T parseJsonList(String raw, Class<T> klass) {
        try {
            return this.mapper.readValue(raw, klass);
        } catch(JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }

    public Optional<Sensor> getByName(String name) {
        Optional<String> rawSensor = this.sapiService.get("/sensor/" + name);
        if(rawSensor.isEmpty()) {
            return Optional.empty();
        }
        return this.parseSensorJson(rawSensor.get());
    }

    private Optional<Sensor> parseSensorJson(String raw) {
        try {
            return Optional.of(this.mapper.readValue(raw, Sensor.class));
        } catch(JsonProcessingException e)  {
            return Optional.empty();
        }
    }
}
