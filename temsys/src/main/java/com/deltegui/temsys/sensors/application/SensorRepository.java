package com.deltegui.temsys.sensors.application;

import com.deltegui.temsys.reports.Report;
import com.deltegui.temsys.sensors.domain.Sensor;

import java.util.List;
import java.util.Optional;

public interface SensorRepository {
    List<Report> getStatusForSensor(String name);
    List<Report> getStatusForAll();
    List<Sensor> getAll();
    Optional<Sensor> getByName(String name);
}
