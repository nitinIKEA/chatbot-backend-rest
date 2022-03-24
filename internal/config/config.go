package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

type dbConfig struct {
	DBHost        string `json:"db_host"`
	DBPort        string `json:"db_port"`
	DBServiceName string `json:"db_service_name"`
	DBUserName    string `json:"db_user_name"`
	DBPassword    string `json:"db_password"`
}

type Conf struct {
	DBConfigDev  dbConfig `json:"dbconfig_dev"`
	DBConfigTest dbConfig `json:"dbconfig_test"`
}

func New(configPath string) (*Conf, error) {
	c := NewDefault()
	if err := c.ReadConfigFromFile(configPath); err != nil {
		log.Fatal(err)
	}
	return c, nil
}

func NewDefault() *Conf {
	return &Conf{}
}

func (c *Conf) ReadConfigFromFile(configPath string) error {
	data, err := ioutil.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil

}
