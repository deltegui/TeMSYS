package cronscheluder

import (
	"fmt"
	"log"
	"temsys"

	"github.com/robfig/cron/v3"
)

// CronScheluder is a implementation of ReportScheluder.
type CronScheluder struct {
	cron *cron.Cron
}

// New ReportScheduler implementation using cron.
func New() temsys.ReportScheluder {
	return CronScheluder{cron.New()}
}

func parseInterval(interval int64) string {
	if interval == 0 {
		return "0"
	}
	return fmt.Sprintf("*/%d", interval)
}

func createCronExpr(interval int64) string {
	minutes := parseInterval(interval % 60)
	return fmt.Sprintf("%s * * * *", minutes)
}

// AddJobEvery interval.
func (scheluder CronScheluder) AddJobEvery(job temsys.ScheluderJob, interval int64) {
	cronExpr := createCronExpr(interval)
	log.Printf("One job have %d inerval, traduced to cron expression: %s\n", interval, cronExpr)
	scheluder.cron.AddFunc(cronExpr, job)
}

// AddJobOnce adds a job that executes only one time
func (scheluder CronScheluder) AddJobOnce(executeJob temsys.ScheluderJob, moment int64) {
	cronExpr := createCronExpr(moment)
	log.Printf("Please execute a refresh configuration job using this cron expression: %s\n", cronExpr)
	var id cron.EntryID
	id, _ = scheluder.cron.AddFunc(cronExpr, func() {
		executeJob()
		scheluder.cron.Remove(id)
	})
}

// Start cron scheduler.
func (scheluder CronScheluder) Start() {
	scheluder.cron.Start()
}

// Stop cron scheduler.
func (scheluder CronScheluder) Stop() {
	log.Println(scheluder.cron.Entries())
	ctx := scheluder.cron.Stop()
	<-ctx.Done()
	log.Println("Stopped Cron scheduler")
	for _, entry := range scheluder.cron.Entries() {
		scheluder.cron.Remove(entry.ID)
	}
	log.Println(scheluder.cron.Entries())
}
