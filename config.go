package wrag

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var config *Config
var endpoints *Endpoints

// Config is a struct containing the configurations of this library
type Config struct {
	Auth struct {
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		UserAgent    string `yaml:"user_agent"`
		AccessToken  string
	}
}

// Endpoints is a struct containing all possible reddit APIs to hit
type Endpoints struct {
	Apis map[string]string
}

func loadConfig(configPath string) *Config {
	data, err := ioutil.ReadFile(configPath)
	check(err)

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	check(err)
	return &config
}

func initialiseConfig(configPath string) {
	config = loadConfig(configPath)
	endpoints = loadEndpoints()
}

func loadEndpoints() *Endpoints {
	data, err := ioutil.ReadFile("./endpoints.yml")
	check(err)

	endpoints := Endpoints{}

	err = yaml.Unmarshal(data, &endpoints)
	check(err)

	return &endpoints
}
