package core

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

type Core struct {
	db database.InterfaceDatabase
}

func NewCore(db database.InterfaceDatabase) *Core {
	return &Core{
		db: db,
	}
}
