package core

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

type Core struct {
	repo database.InterfaceDatabase
}

func NewCore(db database.InterfaceDatabase) *Core {
	return &Core{
		repo: db,
	}
}
