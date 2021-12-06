package pkg

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sidie88/stockbit/app/handler"
	"github.com/sidie88/stockbit/config"
	log "github.com/sirupsen/logrus"
	"time"
)

type EchoWrapper struct {
	Handler    *handler.MovieHandler
	RestServer *echo.Echo
	Config     *config.EchoConfig
	Running    bool
}

func NewEchoWrapper(h *handler.MovieHandler, c *config.EchoConfig) *EchoWrapper {
	return &EchoWrapper{
		Handler:    h,
		RestServer: echo.New(),
		Config:     c,
	}
}

func (e *EchoWrapper) Start() {

	e.RestServer.GET("/search", e.Handler.SearchMovie)
	e.RestServer.GET("/movie-detail", e.Handler.GetMovieDetail)
	go func() {
		log.Info("Echo Server Started at port:", e.Config.Port)
		e.Running = true
		err := e.RestServer.Start(":"+e.Config.Port)
		if err != nil && e.Running {
			log.Error("Failed to start rest api server ", err)
			panic(err)
		}
	}()
}

func (e *EchoWrapper) Stop() {
	e.Running = false
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.RestServer.Shutdown(ctx); err != nil {
		e.RestServer.Logger.Fatal(err)
	}
}
