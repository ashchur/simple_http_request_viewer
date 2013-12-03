package main

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigInfo struct {
	Url     string
	Port    string
	CsUrl   string
	LogFile string
}

var config *ConfigInfo

func Config() *ConfigInfo {
	if config != nil {
		return config
	}
	config = new(ConfigInfo)

	dat, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(dat, config)
	if err != nil {
		panic(err)
	}

	return config
}
