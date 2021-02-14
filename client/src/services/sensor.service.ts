/* eslint-disable no-useless-constructor */
import { store } from '@/store';
import { SensorRepository } from './gateways';

export default class SensorService {
  constructor(private sensorRepo: SensorRepository) {}

  readState(sensorName: string) {
    return this.sensorRepo.getCurrentStateByName(sensorName, store.token?.value ?? '');
  }
}
