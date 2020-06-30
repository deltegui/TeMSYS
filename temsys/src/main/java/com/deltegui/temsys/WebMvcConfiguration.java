package com.deltegui.temsys;

import com.deltegui.temsys.sensors.infraestructure.SensorApi;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class WebMvcConfiguration implements WebMvcConfigurer {

    @Override
    public void addViewControllers(ViewControllerRegistry registry) {
        registry.addViewController("/login").setViewName("session/login");
        registry.addViewController("/error").setViewName("error");
    }

    @Bean
    public SensorApi createSensorApi() {
        return SensorApi.withDefaultClient("http://127.0.0.1:8081");
    }
}
