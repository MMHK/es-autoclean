package main

import (
	"es-autoclean/lib"
	
	"flag"
	"fmt"
)

func main()  {
	confPath := flag.String("c", "conf.json", "config json file")
	flag.Parse()
	
	err, conf := lib.NewConfig(*confPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	running := make(chan bool)
	
	Cron := lib.NewJobScheduler(conf)
	
	Cron.AddJob(conf.CronSpec, func() {
		es := lib.NewESClient(conf)
		list, err := es.ListIndex()
		if err != nil {
			lib.Log.Error(err)
		}
		list = es.FilterExpiredIndex(list)
		if len(list) <= 0 {
			lib.Log.Info("No index will been removed")
		} else {
			for _, l := range list {
				lib.Log.Infof("index(name:%s) will been removed\n", l.Name)
			}
		}
		err = es.RemoveIndex(list)
		if err != nil {
			lib.Log.Error(err)
		}
	})
	
	Cron.Start()
	lib.Log.Info("service started")
	<-running
}