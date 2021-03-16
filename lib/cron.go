package lib

import (
	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	config *Config
	Cron *cron.Cron
}

func NewJobScheduler(conf *Config) *JobScheduler {
	return &JobScheduler{
		config: conf,
		Cron: cron.New(cron.WithSeconds()),
	}
}

func (this *JobScheduler) AddJob(spec string, callback func()) (error)  {
	_, err := this.Cron.AddFunc(spec, callback)
	
	return err
}


func (this *JobScheduler) Start() {
	this.Cron.Start()
}

func (this *JobScheduler) Stop() {
	this.Cron.Stop()
}
