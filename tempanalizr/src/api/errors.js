function isError({ code, reason, fix }) {
  return !!code && !!reason && !!fix;
}

module.exports = {
  isError,
  internal: {
    code: 0,
    reason: 'internal error',
    fix: 'sacrifice a goat',
  },
  invalidRequest: {
    code: 1,
    reason: 'invalid request',
    fix: 'make sure your request is ok',
  },
  sensorAlreadyExists: {
    code: 100,
    reason: 'sensor already exists',
    fix: 'create the sensor',
  },
  sensorNotExists: {
    code: 101,
    reason: 'sensor does not exists',
    fix: 'use the sensor',
  },
  sensorTimeout: {
    code: 200,
    reason: 'sensor does not respond',
    fix: 'check sensor config',
  },
};
