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
app.use((req, _, next) => {
  console.log(`Request from ${req.ip} [${req.method}] ${req.path}`);
  next();
});
api(app);
printLogo();

initialSync()
  .then(startQueueListening)
  .then(() => app.listen(config.port, config.host, () => console.log(`Running on ${config.port}`)))
  .catch(err => {
    console.error(`Error while doing initial sync:`);
    console.error(err);
  });
