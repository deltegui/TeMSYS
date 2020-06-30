package com.deltegui.temsys.sensors.infraestructure;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class ManageSensorsController {
    private final SapiSensorRepository sensorService;

    public ManageSensorsController(SapiSensorRepository sensorService) {
        this.sensorService = sensorService;
    }

    @GetMapping("/dashboard/admin/sensors")
    public String renderManageSensorPage(Model model) {
        var sensors = this.sensorService.getAll();
        model.addAttribute("sensors", sensors);
        return "management/sensor";
    }
}
