using System;
using System.Net.Http;
using System.Threading.Tasks;
using System.Text.Json;
using System.Text;
using System.Collections.Generic;
using System.Linq;
using temsys.Models;

namespace temsys.Services {
    interface IDomainTransformer<T> {
        T ToDomain();
    }

    class SensorConnectionJson {
        public SensorConnectionJson() {
            this.type = "";
            this.value = "";
        }

        public string type { get; set; }
        public string value { get; set; }
    }

    class SensorJson : IDomainTransformer<Sensor> {

        public SensorJson() {
            this.name = "";
            this.connection = new SensorConnectionJson();
            this.updateInterval = 1;
            this.deleted = false;
            this.supportedReports = new string[0];
        }

        public string name { get; set; }
        public SensorConnectionJson connection { get; set; }
        public int updateInterval { get; set; }
        public bool deleted { get; set; }
        public string[] supportedReports { get; set; }

        public Sensor ToDomain() {
            return new Sensor(name, connection.value, updateInterval, supportedReports, deleted);
        }
    }

    class ReportJson : IDomainTransformer<Report> {
        public string type { get; set; }
        public string sensor { get; set; }
        public DateTime date { get; set; }
        public double value { get; set; }

        public Report ToDomain() {
            return new Report(type, sensor, date, value);
        }
    }

    public class SapiRepository : ISensorRepository, IReportTypeRepository {

        private readonly HttpClient client;

        private readonly string baseUrl;

        public SapiRepository(string url) : this(url, new HttpClient()) {}

        public SapiRepository(string url, HttpClient client) {
            this.client = client;
            this.baseUrl = url;
        }

        public async Task<IList<Sensor>> GetAllSensors(bool withDeleted = false) {
            string deletedQuery = (withDeleted) ? "true" : "false";
            return await this.MakeListRequest<Sensor, SensorJson>(client.GetAsync, $"/sensors?deleted={deletedQuery}");
        }

        public async Task<IList<Report>> GetSensorsState() =>
            await this.MakeListRequest<Report, ReportJson>(client.GetAsync, "/sensors/all/now");

        public async Task<Sensor> GetSensorByName(string name) {
            var sensor = await this.MakeRequest<Sensor, SensorJson>(client.GetAsync, $"/sensor/{name}");
            return sensor;
        }

        public async Task<Sensor> SaveSensor(Sensor sensor) {
            string json = JsonSerializer.Serialize(sensor);
            HttpContent content = new StringContent(json, Encoding.UTF8, "application/json");
            var response = await client.PostAsync(FormatUrl("/sensor"), content);
            var rawSensor = await this.SafetlyHandleResponse<SensorJson>(response);
            if(string.IsNullOrEmpty(rawSensor.name)) {
                throw new RepositoryException("Error while communicating with Sapi");
            }
            return rawSensor.ToDomain();
        }

        public async Task DeleteSensorByName(string name) {
            await client.DeleteAsync(FormatUrl($"/sensor/{name}"));
        }

        public async Task<IList<Report>> GetOneSensorStateByName(string name) {
            var url = this.FormatUrl($"/sensor/{name}/now");
            var response = await client.GetAsync(url);
            var handeled = await this.SafetlyHandleResponse<List<ReportJson>>(response);
            return handeled.Select(report => report.ToDomain()).ToList();
        }

        public async Task<IList<string>> GetAllReportTypes() =>
            await this.SafetlyHandleResponse<List<string>>(await client.GetAsync(FormatUrl("/report/types")));

        public async Task SaveReportType(string reportType) {
            await client.PostAsync(FormatUrl($"/report/types/create/{reportType}"), new StringContent(""));
        }

        private async Task<IList<DOMAIN>> MakeListRequest<DOMAIN, TRANSFORMER>(Func<string, Task<HttpResponseMessage>> maker, string endpoint) where TRANSFORMER: IDomainTransformer<DOMAIN> {
            string url = this.FormatUrl(endpoint);
            var response = await maker(url);
            List<TRANSFORMER> raw = await this.SafetlyHandleResponse<List<TRANSFORMER>>(response);
            return raw.Select(s => s.ToDomain()).ToList();
        }

        private async Task<DOMAIN> MakeRequest<DOMAIN, TRANSFORMER>(Func<string, Task<HttpResponseMessage>> maker, string endpoint) where TRANSFORMER : IDomainTransformer<DOMAIN>, new() {
            string url = this.FormatUrl(endpoint);
            var response = await maker(url);
            var raw = await this.SafetlyHandleResponse<TRANSFORMER>(response);
            return raw.ToDomain();
        }

        private string FormatUrl(string endpoint) => $"http://{baseUrl}{endpoint}";

        private async Task<T> SafetlyHandleResponse<T>(HttpResponseMessage response) where T : new() {
            try {
                return await this.HandleResponse<T>(response);
            } catch (Exception e) when (e is HttpRequestException || e is JsonException) {
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
