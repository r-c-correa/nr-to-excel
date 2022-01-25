package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiKey          string `json:"api_key"`
	AccountID       int    `json:"account_id"`
	NRQL            string `json:"nrql"`
	Start           string `json:"start"`
	End             string `json:"end"`
	WindowInMinutes int    `json:"window_in_minutes"`
	PrimaryKey      string `json:"primary_key"`
}

func Load(filename string) (*Config, error) {
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
