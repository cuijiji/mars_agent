package base

import (
	"mars_agent/cmd"
)

var AppConfig Config

type Config struct {
	App struct {
		Env     string `yaml:"env"`
		Port    int    `yaml:"port"`
		Model   string `yaml:"model"`
		baseLog string
	} `yaml:"app"`
	Oss struct {
		EndPoint     string `yaml:"end-point"`
		AccessKey    string `yaml:"access-key"`
		AccessSecret string `yaml:"access-secret"`
	} `yaml:"oss"`
	Commands cmd.Runnable `yaml:"commands"`
}

func (c *Config) IsLocal() bool {
	return c.App.Env == "local"
}

func (c *Config) GetLogDir() string {
	if c.IsLocal() {
		return "./logs/"
	}
	return "/home/gwxdata/wwwlogs/17gwx/" + c.App.Model + "/"

}
