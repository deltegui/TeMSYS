package queues

import (
	"encoding/json"
	"log"
	"sensorapi/src/configuration"
	"sensorapi/src/domain"
	"time"

	"github.com/streadway/amqp"
)

type ReportRabbitMQ struct {
	config  configuration.Configuration
	channel *amqp.Channel
}

func NewReportRabbitMQ(config configuration.Configuration) domain.ReportQueue {
	return &ReportRabbitMQ{
		config,
		nil,
	}
}

func (rabbit *ReportRabbitMQ) Connect() {
	var conn *amqp.Connection
	for {
		connection, err := amqp.Dial(rabbit.config.RabbitMQ)
		if err == nil {
			conn = connection
			break
		}
		log.Println(err)
		time.Sleep(time.Second)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	rabbit.channel = ch
	err = ch.ExchangeDeclare(
		"sapi_reports", // name
		"fanout",       //type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Connected to %s and created the exchange!\n", rabbit.config.RabbitMQ)
}

func (rabbit ReportRabbitMQ) Publish(report domain.Report) {
	data, err := json.Marshal(report)
	if err != nil {
		log.Println("Error while marshaling report: ", err)
		return
	}
	err = rabbit.channel.Publish(
		"sapi_reports", //exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		log.Println("Error while publishing report: ", err)
	}
}
