package domain

import (
	"log"
)

var globalReporter *Reporter

func newScheluderJob(sensor Sensor, reportRepo ReportRepo, queue ReportQueue) ScheluderJob {
	return func() {
		log.Printf("Running job for %s\n", sensor.Name)
		currentReports, err := sensor.GetCurrentState()
		if err != nil {
			log.Println(err)
			return
		}
		for _, report := range currentReports {
			reportRepo.Save(report)
			queue.Publish(report)
		}
	}
}

type Reporter struct {
	sensorRepo  SensorRepo
	reportRepo  ReportRepo
	scheluder   ReportScheluder
	restart     chan bool
	reportQueue ReportQueue
}

func NewReporter(sensorRepo SensorRepo, reportRepo ReportRepo, scheluder ReportScheluder, queue ReportQueue) Reporter {
	if globalReporter == nil {
		log.Println("Created reporter")
		queue.Connect()
		globalReporter = &Reporter{
			sensorRepo:  sensorRepo,
			reportRepo:  reportRepo,
			scheluder:   scheluder,
			restart:     make(chan bool),
			reportQueue: queue,
		}
	}
	return *globalReporter
}

func (reporter Reporter) Start() {
	for {
		sensors, err := reporter.sensorRepo.GetAll(WithoutDeletedSensors)
		if err != nil {
			return
		}
		for _, sensor := range sensors {
			job := newScheluderJob(sensor, reporter.reportRepo, reporter.reportQueue)
			reporter.scheluder.AddJobEvery(job, sensor.UpdateInterval)
		}
		reporter.scheluder.Start()
		<-reporter.restart
		log.Println("Restarting scheduler...")
		reporter.scheluder.Stop()
		log.Println("DONE! Scheduler restarted!")
	}
}

func (reporter Reporter) Restart() {
	reporter.restart <- true
}
