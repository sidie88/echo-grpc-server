package pkg

import (
	"fmt"
	"github.com/sidie88/stockbit/app/service"
	"github.com/sidie88/stockbit/config"
	proto "github.com/sidie88/stockbit/grpc"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"net"
)

type GRpcServerParams struct {
	dig.In
	c  *config.GRpcConfig
	gs *service.GRpcService
}

type GRpcWrapper struct {
	Server   *grpc.Server
	Listener net.Listener
	Port     string
}

func NewGRpcServer(c *config.GRpcConfig, gs *service.GRpcService) (*GRpcWrapper, error) {
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", c.Address, c.Port))
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	proto.RegisterSearchServer(server, gs)

	return &GRpcWrapper{
		Server:   server,
		Listener: list,
		Port:     c.Port,
	}, nil
}

func (g *GRpcWrapper) Start() {

	go func() {
		log.Info("GRpc server started at port:", g.Port)
		err := g.Server.Serve(g.Listener)
		if err != nil {
			log.Error("Failed to start GRpc Servers ", err)
		}
	}()
}
