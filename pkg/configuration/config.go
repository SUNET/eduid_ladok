package configuration

import (
	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"errors"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type envVars struct {
	EduIDConfigYAML string `envconfig:"EDUID_CONFIG_YAML" required:"true"`
}

// Parse parses config file from EDUID_CONFIG_YAML environment variable
func Parse(logger *logger.Logger) (*model.Cfg, error) {
	var env envVars
	logger.Info("Read environmental variable")
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	configPath := env.EduIDConfigYAML

	config := &model.Config{}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, errors.New("config is a folder")
	}

	if err := yaml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	if err := helpers.Check(config, logger); err != nil {
		return nil, err
	}

	return &config.EduID.Worker.Ladok, nil
}
