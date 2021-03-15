package lib

import (
	"testing"
	"time"
)

func TestESClient_ListIndex(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Error(err)
		return
	}
	
	client := NewESClient(conf)
	list, err := client.ListIndex()
	if err != nil {
		t.Error(err)
		return
	}
	
	for _, e := range list {
		t.Log(e.CreateTime().Format(time.RFC3339))
	}
	
	t.Logf("%v", list)
	
	expiredList := client.FilterExpiredIndex(list)
	
	for _, e := range expiredList {
		t.Log(e.CreateTime().Format(time.RFC3339))
	}
	
	t.Logf("%v", expiredList)
}

func TestESClient_RemoveIndex(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Error(err)
		return
	}
	
	client := NewESClient(conf)
	
	err = client.RemoveIndex([]*ESIndex{
		&ESIndex{
			Name: "filebeat-7.3.0-2021.03.02",
		},
	})
	
	if err != nil {
		t.Error(err)
		return
	}
}
