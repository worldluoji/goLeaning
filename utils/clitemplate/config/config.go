package config

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/spf13/viper"

	_ "embed"
)

//go:embed configYaml
var configFiles embed.FS

func init() {

	configBytes, err := configFiles.ReadFile("configYaml/config.yaml")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// fmt.Println(string(configBytes))
	configReader := bytes.NewReader(configBytes)

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(configReader); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

}
