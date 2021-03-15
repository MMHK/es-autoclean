package lib

import (
	"encoding/json"
	"os"
)

type Config struct {
	ESEndPoint    string `json:"es-endpoint"`
	ESIndexPrefix string `json:"index_prefix"`
	KeepDay       int    `json:"keep_day"`
	save_path     string
}

func NewConfig(filename string) (err error, c *Config) {
	c = &Config{}
	c.save_path = filename
	err = c.load(filename)
	return
}

func (c *Config) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (c *Config) Save() error {
	file, err := os.Create(c.save_path)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()
	data, err2 := json.MarshalIndent(c, "", "    ")
	if err2 != nil {
		log.Error(err2)
		return err2
	}
	_, err3 := file.Write(data)
	if err3 != nil {
		log.Error(err3)
	}
	return err3
}
