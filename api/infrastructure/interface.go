package infrastructure

import (
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/Ovsienko023/reporter/infrastructure/database/postgres"
)

type Infrastructure struct {
	Db database.InterfaceDB
}

func New(config configuration.Config) (*Infrastructure, error) {
	db, err := postgres.New(&config.Db)
	if err != nil {
		return nil, err
	}

	return &Infrastructure{
		Db: db,
	}, nil
}

func (i Infrastructure) GetDatabase() database.InterfaceDB {
	return i.Db
}
