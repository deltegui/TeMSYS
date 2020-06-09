const { pool } = require('./connection');

function transformReport({ ID, SENSOR, TYPE, VALUE, REPORT_DATE }) {
  const date = new Date(REPORT_DATE);
  date.setMilliseconds(0);
  return {
    id: ID,
    type: TYPE,
    sensor: SENSOR,
    date,
    value: VALUE,
  };
}

function transformReportResult(arr) {
  return arr.map(transformReport);
}

function formatNumber(number) {
  const str = `0${number}`;
  return str.substr(str.length - 2, str.length);
}

// BE AWARE!! All parameters here must be satinazed.
const transformers = {
  date: {
    transform({ from, to }) {
      if(from === to) {
        const d = new Date(from);
        return `AND REPORTS.REPORT_DATE LIKE "${d.getFullYear()}-${formatNumber(d.getMonth() + 1)}-${formatNumber(d.getDate())}%"`;
      }
      return `AND REPORTS.REPORT_DATE >= "${from}" AND REPORTS.REPORT_DATE <= "${to}"`;
    },
    default: '',
  },
  latestReports: {
    transform(number) {
      return `ORDER BY REPORTS.REPORT_DATE DESC LIMIT ${number}`;
    },
    default: 'ORDER BY REPORTS.REPORT_DATE ASC',
  },
  type: {
    transform(type) {
      return `AND REPORTS.TYPE LIKE "${type}"`;
    },
    default: '',
  },
};

function generateTransformationsFor(options) {
  const transformations = {};
  Object.keys(transformers)
    .forEach(key => {
      if(options[key]) {
        transformations[key] = transformers[key].transform(options[key]);
      } else {
        transformations[key] = transformers[key].default;
      }
    });
  return transformations;
}

const reportRepo = {
  /**
   * Gets all reports for sensor identified by name
   * @param {String} name
   * @returns {Promise<Array>}
   */
  getReportsForSensor(name, options = {}) {
    const fragments = generateTransformationsFor(options);
    return new Promise((resolve, reject) => {
      pool.query(`SELECT
        ID, SENSOR, TYPE, VALUE, REPORT_DATE
      FROM
        REPORTS
      WHERE
        REPORTS.SENSOR = ?
        ${fragments.date}
        ${fragments.type}
      ${fragments.latestReports}`, name, (err, results) => err ? reject(err) : resolve(transformReportResult(results)));
    });
  },

  getRecentReport() {
    return new Promise((resolve, reject) => {
      pool.query(`SELECT
        ID, SENSOR, TYPE, VALUE, REPORT_DATE
      FROM
        REPORTS
      ORDER BY REPORT_DATE DESC
      LIMIT 1`, (err, results) => err ? reject(err) : resolve(transformReportResult(results)[0]));
    });
  },

  save({ type, sensor, date, value }) {
    const values = {
      SENSOR: sensor,
      TYPE: type,
      VALUE: value,
      REPORT_DATE: `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`,
    };
    return new Promise((resolve, reject) => {
      pool.query(`INSERT INTO REPORTS SET ?`, values, err => err ? reject(err) : resolve());
    });
  },
};

module.exports = {
  reportRepo,
};
