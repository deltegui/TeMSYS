package com.deltegui.temsys;

import com.deltegui.temsys.sensors.infraestructure.SensorApi;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class WebMvcConfiguration implements WebMvcConfigurer {

    @Autowired
    private Environment env;

    private final static Logger logger = LoggerFactory.getLogger(WebMvcConfiguration.class);

    @Override
    public void addViewControllers(ViewControllerRegistry registry) {
        registry.addViewController("/login").setViewName("session/login");
        registry.addViewController("/error").setViewName("error");
    }

    @Bean
    public SensorApi createSensorApi() {
        var url = env.getProperty("sapi.url");
        logger.info("USING AS SAPI URL:" + url);
        return SensorApi.withDefaultClient(url);
    }
}
