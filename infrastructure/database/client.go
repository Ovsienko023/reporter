package database

import (
	"context"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	driver Driver
}

type Driver interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close()
}

func New(config *configuration.Db) (*Client, error) {
	pollConf, err := pgxpool.ParseConfig(config.ConnStr)
	if err != nil {
		return nil, err
	}

	pollConf.LazyConnect = true

	pool, err := pgxpool.ConnectConfig(context.Background(), pollConf)
	if err != nil {
		return nil, err
	}

	c := &Client{
		driver: pool,
	}

	return c, nil
}

func (c *Client) Close() {
	c.driver.Close()
}
