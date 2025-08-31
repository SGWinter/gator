package config

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Config struct {
	URL  string `json:"db_url"`
	USER string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(filePath)
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

func (c *Config) SetUser(s string) error {
	c.USER = s

	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg Config) error {
	updatedConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, updatedConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	filePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath = filePath + configLocation
	return filePath, nil
}
