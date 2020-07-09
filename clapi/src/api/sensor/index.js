const { sensorRepo, reportRepo } = require('../../persistence');
const SensorController = require('./sensor.controller');
const ReportService = require('../report/report.service');
const SensorService = require('./sensor.service');

const reportService = new ReportService(reportRepo);
const sensorService = new SensorService(sensorRepo, reportService);
const sensorController = new SensorController(sensorService);

module.exports = sensorController.getMappings.bind(sensorController);
