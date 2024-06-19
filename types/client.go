package types

import "net"

type Client struct {
	Connection net.Conn
	Username   string
}
