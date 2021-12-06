package main

import (
	"github.com/sidie88/stockbit/app"
	"github.com/sidie88/stockbit/app/handler"
	"github.com/sidie88/stockbit/app/service"
	"github.com/sidie88/stockbit/config"
	"github.com/sidie88/stockbit/pkg"
	"go.uber.org/dig"
)

func main() {
	var AppContext = app.ApplicationContext{
		Name:    "Stockbit",
		Configs: app.Configuration{
			&app.Config{
				Prefix: "OMDB",
				Conf:   config.NewOmDbApiConfig,
			},
			&app.Config{
				Prefix: "GRPC",
				Conf:   config.NewGRpcConfig,
			},
			&app.Config{
				Prefix: "ECHO",
				Conf:   config.NewEchoConfig,
			},
		},
		Services: app.Interfaces{
			service.NewOmDbApiService,
			service.NewGRpcService,
			handler.NewMovieHandler,
		},
		Container: dig.New(),
	}

	AppContext.Servers = app.ServerList{
		&app.Server{
			Name:        "GRpc Servers",
			Constructor: pkg.NewGRpcServer,
			StartFunc:   func(g *pkg.GRpcWrapper) { g.Start() },
			StopFunc:    func(g *pkg.GRpcWrapper) { g.Server.GracefulStop() },
		},
		&app.Server{
			Name:        "Echo Servers",
			Constructor: pkg.NewEchoWrapper,
			StartFunc:   func(e *pkg.EchoWrapper) { e.Start() },
			StopFunc:    func(e *pkg.EchoWrapper) { e.Stop() },
		},
	}

	AppContext.Run()
}
