/* eslint-disable class-methods-use-this */

import {
  Sensor,
  Report,
} from '@/services/models';

import {
  SensorRepository,
} from '@/services/gateways';

import makeRequest from './core';

export default class ApiSensorRepository implements SensorRepository {
  async getAll(deleted = false): Promise<Sensor[]> {
    return makeRequest(`/sensors?deleted=${deleted}`);
  }

  async getByName(name: string): Promise<Sensor> {
    return makeRequest(`/sensor/${name}`);
  }

  async getCurrentStateByName(name: string, token: string): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/now`, { method: 'GET', token });
  }

  async getCurrentAverageState(token: string): Promise<Report[]> {
    return makeRequest('/sensors/now/average', { method: 'GET', token });
  }

  async updateSensor(sensor: Sensor, token: string): Promise<Sensor> {
    return makeRequest('/sensor', { method: 'PATCH', body: sensor, token });
  }

  async delete(name: string, token: string): Promise<void> {
    return makeRequest(`/sensor/${name}`, { method: 'DELETE', token });
  }

  async create(sensor: Sensor, token: string): Promise<Sensor> {
    return makeRequest('/sensor', { method: 'POST', body: sensor, token });
  }
}
