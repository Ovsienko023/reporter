package configuration

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Api struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Db struct {
	ConnStr string `yaml:"conn_str"`
}

type Config struct {
	Api Api `yaml:"api"`
	Db  Db  `yaml:"db"`
}

const (
	DefaultConfigPath = ""
	DefaultHost       = "0.0.0.0"
	DefaultPort       = "8888"
	DefaultConnStr    = "postgresql://postgres:1234@database:5432/postgres" // for local debugging use "postgresql://postgres:1234@localhost:5442/postgres"
)

func NewConfig() (*Config, error) {
	cfg := &Config{
		Api{
			Host: DefaultHost,
			Port: DefaultPort,
		},
		Db{
			ConnStr: DefaultConnStr,
		},
	}

	var err error

	switch {
	case *ConfigPathFlag != DefaultConfigPath:
		err = cleanenv.ReadConfig(*ConfigPathFlag, cfg)
	case len(DefaultConfigPath) > 0:
		err = cleanenv.ReadConfig(DefaultConfigPath, cfg)
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
