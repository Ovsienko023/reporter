package database

import (
	"context"
	"github.com/Ovsienko023/reporter/infrastructure/database/dbmessage"
)

type InterfaceDB interface {
	CreateUser(ctx context.Context, msg dbmessage.CreateUser) (*dbmessage.CreatedUser, error)
}
