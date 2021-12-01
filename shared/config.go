package shared

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	StoragePath string `json:"storage_path"`
	CLIConfig CLIConfig `json:"cli"`
}

type CLIConfig struct {
	JenkinsURL string `json:"jenkins_url"`
}

func GetConfig() Config {
	file, err := ioutil.ReadFile("config.json")
	config := Config{}
	err = json.Unmarshal(file, &config)

	if err != nil {
		panic(err)
	}
	return config
}

func SaveConfig(config Config)  {
	jsonBytes, err := json.MarshalIndent(config, "", "\t")
	err = ioutil.WriteFile("config.json", jsonBytes, os.ModePerm) // todo: location

	if err != nil {
		panic(err)
	}
}
