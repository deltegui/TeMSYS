const bodyParser = require('body-parser');
const express = require('express');
const api = require('./api');
const config = require('../config');
const printLogo = require('./logo');
const { initialSync, startQueueListening } = require('./sync.reports');

const app = express();
app.use(express.static('public'));
app.use(bodyParser.json());
app.use((_, res, next) => {
  res.set({
    'Access-Control-Allow-Headers': '*',
    'Access-Control-Allow-Origin': '*',
  });
  next();
});
api(app);
printLogo();

initialSync()
  .then(startQueueListening)
  .then(() => app.listen(config.port, config.host, () => console.log(`Running on ${config.port}`)));
