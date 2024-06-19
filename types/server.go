package types

import "net"

type Server struct {
	Listener        net.Listener
	Clients         []Client
	MaxMessageSize  uint
	MaxUsernameSize uint
	MaxClients      uint
	CurrentClients  uint
}
