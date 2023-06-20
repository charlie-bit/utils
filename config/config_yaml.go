package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Cfg *Config

// Config is the configuration
type Config struct {
	Debug bool `yaml:"debug"`

	Env string `yaml:"-"`

	CronTable []Task `yaml:"cron_table"`
}

type Task struct {
	Name       string `yaml:"name"`
	Timer      string `yaml:"timer"`
	Expiration string `yaml:"expiration"`
	FirstRun   bool   `yaml:"first_run"`
}

// NewConfig creates a Config from file.
func NewConfig(name string) (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	rawData, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("config init failed  {path: %v, err: %v}", name, err.Error())
	}

	out := make(map[string]*Config, 2)
	err = yaml.Unmarshal(rawData, out)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config file failed  {path: %v, err: %v}", name, err.Error())
	}

	c := out[env]
	if c == nil {
		return nil, fmt.Errorf("can't found env config in file  {env: %v, path: %v}", env, name)
	}

	c.Env = env
	Cfg = c
	return c, nil
}
