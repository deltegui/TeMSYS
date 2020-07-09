const errors = require('../errors');
const {
  rejectIfNotExists,
} = require('../helpers');

class SensorService {
  constructor(sensorRepo, reportService) {
    this.sensorRepo = sensorRepo;
    this.reportService = reportService;
  }

  /**
   * Gets a sensor by name
   * @param {String} name
   * @returns {Promise<Sensor>}
   */
  getOneSensorByName(name) {
    return this.sensorRepo.exists(name)
      .then(rejectIfNotExists(errors.sensorNotExists))
      .then(() => this.sensorRepo.getBy(name));
  }

  getSensorStateByName(name) {
    return this.sensorRepo.exists(name)
      .then(rejectIfNotExists(errors.sensorNotExists))
      .then(() => this.sensorRepo.getBy(name))
      .then(() => this.sensorRepo.getCurrentState(name));
  }

  getCurrentStatus() {
    return this.sensorRepo.getAllSensors()
      .then(sensors => Promise.all(sensors
        .map(sensor => this.sensorRepo.getCurrentState(sensor.name).catch(err => {
          console.error(err);
          return {};
        }))))
      .then(reports => reports.flat());
  }

  getAverageCurrentStatus() {
    return this.getCurrentStatus()
      .then(this.reportService.calculateReportAverage);
  }
}

module.exports = SensorService;
