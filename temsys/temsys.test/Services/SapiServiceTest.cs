using System;
using System.Collections.Generic;
using NUnit.Framework;
using temsys.Services;
using Moq;
using System.Net.Http;
using temsys.Models;
using System.Threading;
using System.Threading.Tasks;
using Moq.Protected;

namespace temsys.test.Services {

    [TestFixture]
    class SapiServiceTest {

        private Mock<HttpMessageHandler> httpMessageHandler;

        [SetUp]
        public void initializeMessageHandler() {
            this.httpMessageHandler = new Mock<HttpMessageHandler>();
        }

        public static IEnumerable<object[]> GetAllSensorsData() {
            yield return new object[] { false, new List<Sensor> {
                Sensor.CreateWithDefaults("BombonaSalon", "192.168.1.33"),
                Sensor.CreateWithDefaults("habitacion", "192.168.1.21"),
                Sensor.CreateWithDefaults("salon", "192.168.1.177"),
            }, ResponseMessage(@"
               [
               {
                ""name"": ""BombonaSalon"",
                ""connection"": {
                ""type"": ""http"",
                ""value"": ""192.168.1.33""
                },
                ""updateInterval"": 1,
                ""deleted"": false,
                ""supportedReports"": []
                },
                {
                ""name"": ""habitacion"",
                ""connection"": {
                ""type"": ""http"",
                ""value"": ""192.168.1.21""
                },
                ""updateInterval"": 1,
                ""deleted"": false,
                ""supportedReports"": []
                },
                {
                ""name"": ""salon"",
                ""connection"": {
                ""type"": ""http"",
                ""value"": ""192.168.1.177""
                },
                ""updateInterval"": 1,
                ""deleted"": false,
                ""supportedReports"": []
                }
                ]
            ") };

            yield return new object[] { true, new List<Sensor> { }, ResponseMessage("[]") };

            yield return new object[] { true, new List<Sensor> { }, new HttpResponseMessage(System.Net.HttpStatusCode.BadRequest) };
        }

        [Test, TestCaseSource("GetAllSensorsData")]
        public async Task TestGetAllSensorsShouldReturnAllSensorsInSystem(bool deleted, IList<Sensor> expected, HttpResponseMessage response) {
            var strBool = deleted ? "true" : "false";
            var sapiRepository = this.CreateMockedSapi(response);

            var result = await sapiRepository.GetAllSensors(deleted);

            Assert.AreEqual(expected, result);
            ExpectedCallToEndpoint(HttpMethod.Get, $"/sensors?deleted={strBool}");
        }

        public static IEnumerable<object[]> GetSensorsStateData() {
            yield return new object[] { new List<Report> {
                new Report("temperature", "habitacion", new DateTime(2020, 6, 13, 1, 46, 33), 28.1),
                new Report("temperature", "salon", new DateTime(2020, 6, 13, 1, 46, 33), 27.5),
                new Report("humidity", "salon", new DateTime(2020, 6, 13, 1, 46, 33), 45.2),
            }, ResponseMessage(@"
            [
                {
                    ""type"": ""temperature"",
                    ""sensor"": ""habitacion"",
                    ""date"": ""2020-06-13T01:46:33.0+02:00"",
                    ""value"": 28.1
                },
                {
                    ""type"": ""temperature"",
                    ""sensor"": ""salon"",
                    ""date"": ""2020-06-13T01:46:33.0+02:00"",
                    ""value"": 27.5
                },
                {
                    ""type"": ""humidity"",
                    ""sensor"": ""salon"",
                    ""date"": ""2020-06-13T01:46:33.0+02:00"",
                    ""value"": 45.2
                }
            ]
            ")
            };

            yield return new object[] { new List<Report> { }, ResponseMessage("[]") };

            yield return new object[] { new List<Report> { }, new HttpResponseMessage(System.Net.HttpStatusCode.BadRequest) };
        }

        [Test, TestCaseSource("GetSensorsStateData")]
        public async Task TestGetSensorsStateShouldReturnAllSensorsReports(IList<Report> expected, HttpResponseMessage response) {
            var sapiRepository = this.CreateMockedSapi(response);

            var result = await sapiRepository.GetSensorsState();

            Assert.AreEqual(expected, result);
            ExpectedCallToEndpoint(HttpMethod.Get, $"/sensors/all/now");
        }

        public static IEnumerable<object[]> GetSensorByNameData() {
            yield return new object[] { "habitacion", Sensor.CreateWithDefaults("habitacion", "192.168.1.21"), ResponseMessage(@"
            {
                ""name"": ""habitacion"",
                ""connection"": {
                    ""type"": ""http"",
                    ""value"": ""192.168.1.21""
                },
                ""updateInterval"": 1,
                ""deleted"": false,
                ""supportedReports"": []
            }
            ") };

            yield return new object[] { "salon", Sensor.CreateWithDefaults("", ""), ResponseMessage("") };

            yield return new object[] { "salon", Sensor.CreateWithDefaults("", ""), new HttpResponseMessage(System.Net.HttpStatusCode.BadRequest) };
        }

        [Test, TestCaseSource("GetSensorByNameData")]
        public async Task TestGetSensorByNameShouldReturnOneSensorOrEmtpyOne(string sensorName, Sensor expected, HttpResponseMessage response) {
            var sapiRepository = this.CreateMockedSapi(response);

            var result = await sapiRepository.GetSensorByName(sensorName);

            Assert.AreEqual(expected, result);
            ExpectedCallToEndpoint(HttpMethod.Get, $"/sensor/{sensorName}");
        }

        public static IEnumerable<object[]> SaveSensorData() {
            yield return new object[] { Sensor.CreateWithDefaults("prueba", "192.168.1.1"), ResponseMessage(@"
            {
                ""name"": ""prueba"",
                ""connection"": {
                    ""type"": ""http"",
                    ""value"": ""192.168.1.1""
                },
                ""updateInterval"": 1,
                ""deleted"": false,
                ""supportedReports"": []
            }
            ") };
        }

        [Test, TestCaseSource("SaveSensorData")]
        public async Task TestSaveSensorShouldSendTheSensorToSaveIt(Sensor sensorToSave, HttpResponseMessage response) {
            var sapiRepository = this.CreateMockedSapi(response);
            var result = await sapiRepository.SaveSensor(sensorToSave);
            Assert.AreEqual(sensorToSave, result);
            ExpectedCallToEndpoint(HttpMethod.Post, $"/sensor");
        }

        [Test]
        public void TestSaveSensorShouldRaiseExceptionIfSapiFails() {
            var sensor = Sensor.CreateWithDefaults("prueba", "192.168.1.1");
            var response = ResponseMessage(@"
            {
                ""Code"": 102,
                ""Reason"": ""Sensor already exists"",
                ""Fix"": ""Use the sensor""
            }");
            var sapiRepository = this.CreateMockedSapi(response);
            Assert.ThrowsAsync<RepositoryException>(() => sapiRepository.SaveSensor(sensor));
            ExpectedCallToEndpoint(HttpMethod.Post, $"/sensor");
        }

        [Test]
        public async Task TestDeleteSensorByNameShouldSendSensorDelete() {
            var sensorName = "habitacion";
            var sapiRepository = this.CreateMockedSapi(ResponseMessage(""));
            await sapiRepository.DeleteSensorByName(sensorName);
            ExpectedCallToEndpoint(HttpMethod.Delete, $"/sensor/{sensorName}");
        }

        public static IEnumerable<object[]> GetOneSensorStateByNameData() {
            yield return new object[] { "salon", new List<Report> {
                new Report("temperature", "salon", new DateTime(2020, 6, 13, 1, 46, 33), 27.5),
                new Report("humidity", "salon", new DateTime(2020, 6, 13, 1, 46, 33), 45.2),
            }, ResponseMessage(@"
            [
                {
                    ""type"": ""temperature"",
                    ""sensor"": ""salon"",
                    ""date"": ""2020-06-13T01:46:33.0+02:00"",
                    ""value"": 27.5
                },
                {
                    ""type"": ""humidity"",
                    ""sensor"": ""salon"",
                    ""date"": ""2020-06-13T01:46:33.0+02:00"",
                    ""value"": 45.2
                }
            ]
            ")
            };

            yield return new object[] { "salon", new List<Report> { }, ResponseMessage("[]") };

            yield return new object[] { "salon", new List<Report> { }, new HttpResponseMessage(System.Net.HttpStatusCode.BadRequest) };
        }

        [Test, TestCaseSource("GetOneSensorStateByNameData")]
        public async Task TestGetOneSensorStateByNameDateShouldReturnCurrentState(string sensorName, IList<Report> expected, HttpResponseMessage response) {
            var httpMessageHandler = new Mock<HttpMessageHandler>();
            var sapiRepository = this.CreateMockedSapi(response);

            var result = await sapiRepository.GetOneSensorStateByName(sensorName);

            Assert.AreEqual(expected, result);
            ExpectedCallToEndpoint(HttpMethod.Get, $"/sensor/{sensorName}/now");
        }

        [Test]
        public async Task TestGetAllReportTypesMustReturnAllReportTypesDefined() {
            var expectedTypes = new List<string> { "humidity", "temperature" };
            var httpMessageHandler = new Mock<HttpMessageHandler>();
            var sapiRepository = this.CreateMockedSapi(ResponseMessage(@"
            [
                ""humidity"",
                ""temperature""
            ]"));
            var reportTypes = await sapiRepository.GetAllReportTypes();
            Assert.AreEqual(expectedTypes, reportTypes);
            ExpectedCallToEndpoint(HttpMethod.Get, "/report/types");
        }

        [Test]
        public async Task TestSaveReportTypeShouldCreateNewReportType() {
            var reportType = "hola";
            var sapiRepository = this.CreateMockedSapi(ResponseMessage(@"
            {
                ""ReportType"": ""hola""
            }"));
            await sapiRepository.SaveReportType(reportType);
            ExpectedCallToEndpoint(HttpMethod.Post, $"/report/types/create/{reportType}");
        }


        private SapiRepository CreateMockedSapi(HttpResponseMessage response) {
            httpMessageHandler
                .Protected()
                .Setup<Task<HttpResponseMessage>>("SendAsync", ItExpr.IsAny<HttpRequestMessage>(), ItExpr.IsAny<CancellationToken>())
                .ReturnsAsync(response)
                .Verifiable();
            return new SapiRepository("localhost:8080", new HttpClient(httpMessageHandler.Object));
        }

        private void ExpectedCallToEndpoint(HttpMethod method, string endpoint) {
            var expectedUri = new Uri($"http://localhost:8080{endpoint}");
            httpMessageHandler.Protected().Verify(
               "SendAsync",
               Times.Once(),
               ItExpr.Is<HttpRequestMessage>(req =>
                  req.Method == method
                  && req.RequestUri == expectedUri
               ),
               ItExpr.IsAny<CancellationToken>()
            );
        }

        public static HttpResponseMessage ResponseMessage(string rawJson) {
            return new HttpResponseMessage {
                Content = new StringContent(rawJson),
                StatusCode = System.Net.HttpStatusCode.OK,
            };
        }
    }
}
