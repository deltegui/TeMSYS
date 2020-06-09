import config from '../../config';
import {
  SensorApi,
  ReportApi,
} from './api';

const apiCore = {
  makeRequest(endpoint, body = undefined) {
    let reqConfig = undefined;
    if(body) {
      reqConfig = {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
          'Content-Type': 'application/json',
        },
      };
    }
    return fetch(`${config.apiURL}${endpoint}`, reqConfig)
      .then(res => res.json());
  }
};

export default {
  sensor: new SensorApi(apiCore),
  report: new ReportApi(apiCore),
};