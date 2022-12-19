package configuration

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Doc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Api struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Doc  Doc    `yaml:"doc"`
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

	DefaultApiHost = "0.0.0.0"
	DefaultApiPort = "8888"

	DefaultDocHost = "85.193.83.76"
	//DefaultDocHost = "localhost"
	DefaultDocPort = "8888"

	//DefaultDbConnStr = "postgresql://postgres:1234@localhost:5442/postgres"
	DefaultDbConnStr = "postgresql://postgres:1234@database:5432/postgres"
)

func NewConfig() (*Config, error) {
	cfg := &Config{
		Api{
			Host: DefaultApiHost,
			Port: DefaultApiPort,
			Doc: Doc{
				Host: DefaultDocHost,
				Port: DefaultDocPort,
			},
		},
		Db{
			ConnStr: DefaultDbConnStr,
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
