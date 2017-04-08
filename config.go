package ursh

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config is a configuration struct
type Config struct {
	Storage struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	ExpireTime int `yaml:"expire_time"`
	UIPort     int
}

// Parse parses the config file
func Parse(filename string) *Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read config file %s: %s", filename, err)
	}
	c := Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("Unable to parse config file %s: %s", filename, err)
	}
	return &c
}
