package database

type Client struct {
	connStr *string
	driver  Driver
}

type Driver struct {
	conn interface{}
}

func NewClient(connStr *string) (*Client, error) {
	return &Client{}, nil
}
