using System;
using System.Net.Http;
using System.Threading.Tasks;
using temsys.Sensors.Domain;
using System.Text.Json;
using System.Collections.Generic;

namespace temsys.Sensors.Implementation {
    public class HttpSapiService : ISensorRepository, IReportTypeRepository {

        private static readonly HttpClient client = new HttpClient();

        private string baseUrl;

        public HttpSapiService(string url) {
            this.baseUrl = url;
        }

        public async Task<IList<Sensor>> GetAllSensors(bool withDeleted = false) =>
            await this.SafetlyHandleResponse<List<Sensor>>(await client.GetAsync(""));

        public async Task<IList<Report>> GetSensorsState() =>
            await this.SafetlyHandleResponse<List<Report>>(await client.GetAsync(""));

        public async Task<Sensor> GetSensorByName(string name) {
            return new Sensor();
        }

        public void SaveSensor(Sensor sensor) {
        }

        public void DeleteSensorByName(string name) {
        }

        public Report[] GetOneSensorStateByName(string name) {
            return null;
        }

        public string[] GetAllReportTypes() {
            return null;
        }

        public void SaveReportType(string reportType) {
        }

        private async Task<T> SafetlyHandleResponse<T>(HttpResponseMessage response) where T: new() {
            try {
                return await this.HandleResponse<T>(response);
            } catch(HttpRequestException) {
                return new T();
            }
        }

        private async Task<T> HandleResponse<T>(HttpResponseMessage response) {
            response.EnsureSuccessStatusCode();
            byte[] responseBody = await response.Content.ReadAsByteArrayAsync();
            return JsonSerializer.Deserialize<T>(responseBody);
        }

    }
}