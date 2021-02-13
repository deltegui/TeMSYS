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
  async getAll(): Promise<Sensor[]> {
    return makeRequest('/sensors');
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
}
