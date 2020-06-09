const amqp = require('amqplib/callback_api');
const { reportRepo, sensorRepo } = require("./persistence");
const config = require('../config.json');

async function initialSync() {
  console.log("Synchronizing with SensorAPI");
  let report = await reportRepo.getRecentReport();
  if(!report) {
    report = {
      date: new Date(1970),
    };
  }
  report.date.setSeconds(report.date.getSeconds() + 1);
  const missingReports = await sensorRepo.getReportsBetweenDates(report.date, new Date());
  missingReports.forEach(rep => {
    const r = rep;
    process.stdout.write('.');
    r.date = new Date(r.date);
    reportRepo.save(r);
  });
}

function handleRabbitMsg(msg) {
  const report = JSON.parse(msg.content);
  console.log('Recieved report from rabbit: ', report);
  report.date = new Date(report.date);
  reportRepo.save(report);
}

async function startQueueListening() {
  amqp.connect(config.rabbitmq.url, (connErr, conn) => {
    if(connErr) {
      console.error(`Error while connecting to RabbitMQ: ${connErr}`);
      process.exit(2);
    }
    conn.createChannel((chanErr, channel) => {
      if(chanErr) {
        console.error(`Error creating channel in RabbitMQ: ${chanErr}`);
        process.exit(2);
      }
      channel.assertExchange(config.rabbitmq.exchange, 'fanout', { durable: true });
      channel.assertQueue('', { exclusive: true }, (queueErr, queue) => {
        if(queueErr) {
          console.error(`Error binding channel into exchange ${config.rabbitmq.exchange}: ${queueErr}`);
          process.exit(2);
        }
        channel.bindQueue(queue.queue, config.rabbitmq.exchange, '');
        channel.consume(queue.queue, handleRabbitMsg);
      });
    });
  });
}

module.exports = {
  initialSync,
  startQueueListening,
};
