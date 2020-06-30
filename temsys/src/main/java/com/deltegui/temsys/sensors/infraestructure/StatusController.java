package com.deltegui.temsys.sensors.infraestructure;

import com.deltegui.temsys.sensors.application.SensorService;
import com.deltegui.temsys.sensors.application.SensorStatusModel;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.List;

@Controller
public class StatusController {
    private final SensorService sensorService;

    public StatusController(SensorService sensorService) {
        this.sensorService = sensorService;
    }

    @GetMapping("/dashboard/status")
    public String showCurrentSystemStatus(Model model) {
        List<SensorStatusModel> sensors = this.sensorService.getAllSensorsWithStatus();
        model.addAttribute("sensors", sensors);
        return "dashboard/status";
    }

}
