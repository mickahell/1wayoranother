package helpers

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed conf.yaml
var conf_file []byte

type Cfg struct {
	RouterPass  string `yaml:"router_pass"`
	IsConnected bool   `yaml:"is_connected"`
}

var AppConfig Cfg

func ReadConfig() {
	err := yaml.Unmarshal(conf_file, &AppConfig)
	if err != nil {
		fmt.Println(err)
	}
}

func TheAppConfig() *Cfg {
	return &AppConfig
}
