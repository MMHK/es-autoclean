package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	es6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type ESClient struct {
	conf *Config
}

type ESIndex struct {
	Name        string `json:"index"`
	Status      string `json:"status"`
	Health      string `json:"health"`
	CreatedTime string `json:"creation.date"`
}

func (this *ESIndex) CreateTime() time.Time {
	raw := strings.Trim(this.CreatedTime, `"`)
	millis, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		millis = 0;
		log.Error(err)
	}
	return time.Unix(0, millis * int64(time.Millisecond));
}

func NewESClient(c *Config) *ESClient {
	return &ESClient{
		conf: c,
	}
}

func (this *ESClient) ListIndex() ([]*ESIndex, error) {
	list := make([]*ESIndex, 0)
	
	client, err := es6.NewClient(es6.Config{
		Addresses: []string{this.conf.ESEndPoint},
	})
	if err != nil {
		log.Error(err)
		return list, err
	}
	req := esapi.CatIndicesRequest{
		Index: []string{fmt.Sprintf(`%s*`, this.conf.ESIndexPrefix)},
		Format: "json",
		H: []string{"index","status","health","creation.date"},
	}
	rep, err := req.Do(context.Background(), client)
	if err != nil {
		log.Error(err)
		return list, err
	}
	decoder := json.NewDecoder(rep.Body)
	err = decoder.Decode(&list)
	if err != nil {
		log.Error(err)
		return list, err
	}
	
	return list, nil
}

func (this *ESClient) FilterExpiredIndex(list []*ESIndex) []*ESIndex {
	out := make([]*ESIndex, 0)
	
	now := time.Now().Add( time.Duration(-24 * this.conf.KeepDay) * time.Hour)
	
	for _, e := range list {
		if now.After(e.CreateTime()) {
			out = append(out, e)
		}
	}
	
	return out
}

func (this *ESClient) RemoveIndex(list []*ESIndex) error {
	client, err := es6.NewClient(es6.Config{
		Addresses: []string{this.conf.ESEndPoint},
	})
	if err != nil {
		log.Error(err)
		return err
	}
	
	nameList := []string{}
	
	for _, e := range list {
		nameList = append(nameList, e.Name)
	}
	
	req := esapi.IndicesDeleteRequest{
		Index: nameList,
	}
	resp, err := req.Do(context.Background(), client)
	if err != nil {
		log.Error(err)
		return err
	}
	
	if resp.IsError() {
		bin, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}
		errorText := string(bin)
		log.Error(errorText)
		return errors.New(errorText)
	}
	
	return nil
}