package com.deltegui.temsys.sensors.infraestructure;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;
import java.util.Optional;

public class SensorApi {
    private final HttpClient client;
    private final String baseUrl;

    public SensorApi(HttpClient client, String url) {
        this.baseUrl = url;
        this.client = client;
    }

    public static SensorApi withDefaultClient(String url) {
        var client = HttpClient.newBuilder()
                .connectTimeout(Duration.ofSeconds(1))
                .build();
        return new SensorApi(client, url);
    }

    public Optional<String> get(String endpoint) {
        return this.makeRequest(this.createRequest(endpoint).build());
    }

    public Optional<String> post(String endpoint, String rawPayload) {
        return this.makeRequest(this.createRequest(endpoint)
                .POST(HttpRequest.BodyPublishers.ofString(rawPayload))
                .build());
    }

    public Optional<String> delete(String endpoint) {
        return this.makeRequest(this.createRequest(endpoint)
                .DELETE()
                .build());
    }

    private HttpRequest.Builder createRequest(String endpoint) {
        try {
            return HttpRequest.newBuilder()
                .uri(new URI(this.baseUrl + endpoint));
        } catch(URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    private Optional<String> makeRequest(HttpRequest request) {
        try {
            var response = this.client.send(request, HttpResponse.BodyHandlers.ofString());
            return Optional.of(response.body());
        } catch (IOException | InterruptedException e) {
            return Optional.empty();
        }
    }
}
