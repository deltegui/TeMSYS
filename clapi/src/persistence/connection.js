const mysql = require('mysql');
const { database } = require('../../config');

mysql
  .createConnection(database)
  .connect(err => {
    if(err) {
      console.error(`MYSQL: cannot connect to database: ${database.host}::${database.database}`);
      console.error(err);
      process.exit(1);
    }
    console.log('MYSQL: connected to database');
  });

const pool = mysql.createPool(database);
pool.on('error', console.error);

module.exports = {
  pool,
};
