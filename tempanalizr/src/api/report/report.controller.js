const validator = require('simplejsonvalidator')();
const errors = require('../errors');
const { handleErr } = require('../helpers');

validator.create(t => ({
  date: {
    from: t.date,
    to: t.date,
  },
  latestReports: t.number.positive,
  average: t.boolean,
  type: t.string,
}), 'reportsWithFilter');

class ReportController {
  constructor(reportService) {
    this.reportService = reportService;
  }

  getReportsForSensor(req, res) {
    if(!req.params.name) {
      res.send(errors.invalidRequest);
      return;
    }
    this.reportService.getReportsForSensor(req.params.name)
      .then(res.send.bind(res))
      .catch(handleErr(res.send.bind(res)));
  }

  getReportsForSensorWithFilters(req, res) {
    const options = req.body;
    if(!validator.validate(options, 'reportsWithFilter')) {
      res.send(errors.invalidRequest);
      return;
    }
    if(!req.params.name) {
      res.send(errors.invalidRequest);
      return;
    }
    this.reportService.getReportsForSensor(req.params.name, options)
      .then(res.send.bind(res))
      .catch(handleErr(res.send.bind(res)));
  }

  getMappings(router) {
    router.get('/:name/reports', this.getReportsForSensor.bind(this));
    router.post('/:name/reports', this.getReportsForSensorWithFilters.bind(this));
  }
}

module.exports = ReportController;
