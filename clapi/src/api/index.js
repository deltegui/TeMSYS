/* eslint-disable global-require */

const express = require('express');

const routerCreators = [
  { url: '/api/sensors', mapper: require('./sensors') },
  { url: '/api/sensor', mapper: require('./sensor') },
  { url: '/api/sensor', mapper: require('./report') },
];

module.exports = app => routerCreators.forEach(({ url, mapper }) => {
  const r = express.Router();
  mapper(r);
  app.use(url, r);
});
