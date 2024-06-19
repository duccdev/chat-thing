package server

import (
	"chat-thing/types"
	"fmt"
	"net"
)

func CreateServer(addr string, port uint16) (types.Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	maxClients := 3
	maxMessageSize := 4000
	maxUsernameSize := 16

	if err != nil {
		return types.Server{}, err
	}

	return types.Server{
		Listener:        listener,
		Clients:         make([]types.Client, maxClients),
		MaxMessageSize:  uint(maxMessageSize),
		MaxUsernameSize: uint(maxUsernameSize),
	}, nil
}

func Broadcast(server types.Server, msg string) {
	for _, client := range server.Clients {
		client.Connection.Write([]byte(msg))
	}
}
