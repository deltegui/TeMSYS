// Persistence using SensorAPI
const http = require('http');
const { sensorAPI } = require('../../config.json');

const HTTP_GET = 'GET';
const HTTP_POST = 'POST';

function makeRequest({ method, body, path }) {
  return new Promise((resolve, reject) => {
    const data = JSON.stringify(body);
    const options = {
      hostname: sensorAPI.hostname,
      port: sensorAPI.port,
      path,
      method,
      timeout: 3000,
    };
    if(method === HTTP_POST) {
      options.headers = {
        'Content-Type': 'application/json',
        'Content-Length': data.length,
      };
    }
    const req = http.request(options, resp => {
      let buffer = '';
      resp.on('data', chunk => {
        buffer += chunk;
      });
      resp.on('end', () => {
        const response = JSON.parse(buffer);
        if(response.code) {
          reject(response);
        } else {
          resolve(response);
        }
      });
    });
    req.on('error', reject);
    if(method === HTTP_POST) req.write(data);
    req.end();
  });
}

const sensorRepo = {
  /**
   * Gets all sensors
   * @returns {Promise<[Sensor]>} returns a sensor array or error.
   */
  getAllSensors() {
    return makeRequest({
      method: HTTP_GET,
      path: '/sensors',
    });
  },

  /**
   * Get a sensor by name
   * @param {String} name
   * @returns {Promise<Sensor>} returns a sensor or error.
   */
  getBy(name) {
    return makeRequest({
      method: HTTP_GET,
      path: `/sensor/${name}`,
    });
  },

  /**
   * Check if a sensor exists with name
   * @param {String} name
   * @returns {Promise<Boolean>}
   */
  exists(name) {
    return this.getBy(name).then(() => true, () => false);
  },

  /**
   * Gets current status for sensor identified by name
   * @param {String} name
   * @returns {Promise<Array<Report>>}
   */
  getCurrentState(name) {
    return makeRequest({
      method: HTTP_GET,
      path: `/sensor/${name}/now`,
    });
  },

  getReportsBetweenDates(from, to) {
    return makeRequest({
      method: HTTP_GET,
      path: `/report/dates?from=${from.toJSON()}&to=${to.toJSON()}`,
    });
  },
};

module.exports = { sensorRepo };
