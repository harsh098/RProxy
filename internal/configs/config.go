package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type resource struct {
	Name         string `yaml:"name"`
	Endpoint     string `yaml:"endpoint"`
	Upstream_URL string `yaml:"upstream_url"`
}

type configuration struct {
	Gateway struct {
		Host        string `yaml:"host"`
		Listen_port string `yaml:"listen_port"`
		Scheme      string `yaml:"scheme"`
	} `yaml:"gateway"`

	Resources []resource `yaml:"resources"`
}

var Config *configuration

func NewConfig() (*configuration, error) {
	viper.AddConfigPath("data")
	viper.AddConfigPath("/etc/RPServer/data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("[FATAL] Error loading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("[FATAL] Cannot read config file: %s", err)
	}

	return Config, nil
}
