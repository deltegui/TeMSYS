import ApiSensorRepository from '@/impl/api/sensor.repo';
import ApiReportRepository from '@/impl/api/report.repo';
import ApiUserRepository from '@/impl/api/user.repo';
import LocalStorageTokenRepository from '@/impl/tokenstorage';
import UserService from './user.service';
import ReportService from './report.service';
import SensorService from './sensor.service';

const sensorRepo = new ApiSensorRepository();
const reportRepo = new ApiReportRepository();
const userRepo = new ApiUserRepository();
const storageTokenRepo = new LocalStorageTokenRepository();

export const userService = new UserService(userRepo, storageTokenRepo);
export const reportService = new ReportService(reportRepo, sensorRepo);
export const sensorService = new SensorService(sensorRepo);
