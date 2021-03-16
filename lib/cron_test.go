package lib

import (
	"testing"
	"time"
)

func TestJobScheduler_AddJob(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Error(err)
		return
	}
	
	cron := NewJobScheduler(conf)
	cron.AddJob("0/1 * * * * ?", func() {
		Log.Info("喵~~~~~")
	})
	
	queue := make(chan bool, 0)
	defer close(queue)
	
	time.AfterFunc(time.Second * 10, func() {
		cron.Stop()
		queue<-true
	})
	
	cron.Start()
	
	<-queue
}