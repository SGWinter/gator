package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	URL string `json:"db_url"`
}

func Read() (Config, error) {
	dat, err := os.ReadFile(configLocation)
	if err != nil {
		return Config{}, err
	}

	configRead := Config{}
	err = json.Unmarshal(dat, &configRead)
	if err != nil {
		return Config{}, err
	}

	return configRead, nil
}
