package configuration

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not found locale .env file", err.Error())
	}
}

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

	DefaultDocHost = "127.0.0.1"
	DefaultDocPort = "8888"

	DefaultDbConnStr = "postgresql://postgres:1234@localhost:5442/postgres"
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

	if dbConn, ok := os.LookupEnv("RP_DATABASE_CONN_STRING"); ok {
		cfg.Db.ConnStr = dbConn
	}

	if docHost, ok := os.LookupEnv("RP_DOC_HOST"); ok {
		cfg.Api.Doc.Host = docHost
	}

	if docPort, ok := os.LookupEnv("RP_DOC_PORT"); ok {
		cfg.Api.Doc.Port = docPort
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
