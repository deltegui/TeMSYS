/* eslint-disable no-useless-constructor */
import { store } from '@/store';
import { SensorRepository } from './gateways';
import { Sensor } from './models';

export default class SensorService {
  constructor(private sensorRepo: SensorRepository) {}

  readState(sensorName: string) {
    return this.sensorRepo.getCurrentStateByName(sensorName, store.token?.value ?? '');
  }

  async getAll(): Promise<Sensor[]> {
    return this.sensorRepo.getAll(false);
  }

  async getAllWithDeleted(): Promise<Sensor[]> {
    return this.sensorRepo.getAll(true);
  }

  async getOne(name: string): Promise<Sensor> {
    return this.sensorRepo.getByName(name);
  }

  async update(sensor: Sensor): Promise<Sensor> {
    return this.sensorRepo.updateSensor(sensor, store.token?.value ?? '');
  }

  async delete(name: string): Promise<void> {
    return this.sensorRepo.delete(name, store.token?.value ?? '');
  }

  async create(sensor: Sensor): Promise<Sensor> {
    return this.sensorRepo.create(sensor, store.token?.value ?? '');
  }
}
