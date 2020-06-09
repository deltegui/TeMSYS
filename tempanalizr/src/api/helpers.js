const errors = require('./errors');

module.exports = {
  handleErr(send) {
    return err => {
      console.error(err);
      if(errors.isError(err)) {
        send(err);
      } else {
        send(errors.internal);
      }
    };
  },

  rejectIfExists(error) {
    return exists => exists ? Promise.reject(error) : Promise.resolve();
  },

  rejectIfNotExists(error) {
    return exists => exists ? Promise.resolve() : Promise.reject(error);
  },
};
