package config

import (
	"common/database"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type AreaServerConfig struct {
	Port string `yaml:"port"`
}
type AreaSchedulerConfig struct {
	Path string `yaml:"path"`
}

type AreaConfig struct {
	Server    AreaServerConfig    `yaml:"server"`
	Scheduler AreaSchedulerConfig `yaml:"scheduler"`
	Database  database.DBConfig   `yaml:"database"`
}

func (c *AreaConfig) LoadConfig(file string) error {
	files, err := os.ReadFile(file)
	if err != nil {
		logrus.WithError(err).Error("config file read failure")
		return err
	}

	if err = yaml.Unmarshal(files, c); err != nil {
		return err
	}

	return nil
}
