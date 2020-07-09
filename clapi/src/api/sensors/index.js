const { sensorRepo, reportRepo } = require('../../persistence');
const ReportService = require('../report/report.service');
const SensorService = require('../sensor/sensor.service');

const reportService = new ReportService(reportRepo);
const sensorService = new SensorService(sensorRepo, reportService);
const errors = require('../errors');

function getAllSensors(_, res) {
  sensorRepo.getAllSensors()
    .then(sensors => res.json(sensors))
    .catch(() => res.send(errors.internal));
}

function getAverageCurrentStatus(_, res) {
  sensorService.getAverageCurrentStatus()
    .then(sensors => res.json(sensors))
    .catch(() => res.send(errors.internal));
}

function getCurrentStatus(_, res) {
  sensorService.getCurrentStatus()
    .then(sensors => res.json(sensors))
    .catch(() => res.send(errors.internal));
}

module.exports = router => {
  router.get('/', getAllSensors);
  router.get('/now/average', getAverageCurrentStatus);
  router.get('/now', getCurrentStatus);
};
