const { reportRepo } = require('../../persistence');
const ReportService = require('./report.service');
const ReportController = require('./report.controller');

const reportService = new ReportService(reportRepo);
const reportController = new ReportController(reportService);

module.exports = reportController.getMappings.bind(reportController);
