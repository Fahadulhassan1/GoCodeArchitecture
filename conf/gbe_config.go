package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var (
	configPath     = "."
	configFileName = "conf.json"
)

type GbeConfig struct {
	RestServer RestServerConfig `json:"restServer"`
}
type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config GbeConfig
var configOnce sync.Once

func SetConfFilePath(path string) {
	configPath = path
}

func SetConfFileName(name string) {
	configFileName = name
}

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile(configPath + "/" + configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
