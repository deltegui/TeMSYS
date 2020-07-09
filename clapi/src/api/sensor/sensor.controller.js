const errors = require('../errors');
const { handleErr } = require('../helpers');

// REGEX FOR IPs
// /^((([0-1][0-9][0-9]|2[0-4][0-9]|25[0-5])|\d{1,2})\.){3}(([0-1][0-9][0-9]|2[0-4][0-9]|25[0-5])|\d{1,2})$/;

class SensorController {
  constructor(sensorService) {
    this.sensorService = sensorService;
  }

  getOneByName(req, res) {
    if(!req.params.name) {
      res.send(errors.invalidRequest);
      return;
    }
    this.sensorService.getOneSensorByName(req.params.name)
      .then(res.send.bind(res))
      .catch(handleErr(res.send.bind(res)));
  }

  getCurrentSensorState(req, res) {
    if(!req.params.name) {
      res.send(errors.invalidRequest);
      return;
    }
    this.sensorService.getSensorStateByName(req.params.name)
      .then(res.send.bind(res))
      .catch(handleErr(res.send.bind(res)));
  }

  getMappings(router) {
    router.get('/:name', this.getOneByName.bind(this));
    router.get('/:name/now', this.getCurrentSensorState.bind(this));
  }
}

module.exports = SensorController;
