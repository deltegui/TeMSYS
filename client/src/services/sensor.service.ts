/* eslint-disable no-useless-constructor */

import { SensorRepository } from './gateways';

export default class SensorService {
  constructor(private sensorRepo: SensorRepository) {}

  readState(sensorName: string, token: string) {
    return this.sensorRepo.getCurrentStateByName(sensorName, token);
  }
}
