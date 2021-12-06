package app

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Interfaces is slice of interface{}
type Interfaces []interface{}

type Config struct {
	Prefix string
	Conf   interface{}
}

type Configuration []*Config

type Server struct {
	Name        string
	Constructor interface{}
	StartFunc   interface{}
	StopFunc    interface{}
}

type ServerList []*Server

type ApplicationContext struct {
	Name       string
	Configs    Configuration
	configFunc Interfaces
	Services   Interfaces
	Servers    ServerList
	Container  *dig.Container
}

func (a *ApplicationContext) Run() {
	log.Info("Starting Application:", a.Name)
	defer log.Info("Shutdown Application:", a.Name)

	a.provideConfig()
	a.provideServices()

	wg := sync.WaitGroup{}
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)



	go func() {
		stop := <-gracefulStop
		log.Info("Stopping all services by ", stop)
		for _, s := range a.Servers {

			err := a.Container.Invoke(s.StopFunc)
			if err != nil {
				log.Errorf("Failed to gracefully shutdown %s\n Error: %+v", s.Name, err)
			} else {
				log.Info("Gracefully shutdown ", s.Name)
			}
			wg.Done()
		}
	}()

	for _, s := range a.Servers {
		wg.Add(1)
		log.Info("Constructing Servers:", s.Name)
		err := a.Container.Provide(s.Constructor)
		if err != nil {
			log.Error("Failed to construct:", s.Name)
			panic(err)
		}
	}

	for _, sv := range a.Servers {
		startServer(a, sv)
	}

	wg.Wait()
}

func startServer(a *ApplicationContext, sv *Server) {
	log.Info("Starting Servers:", sv.Name)
	err := a.Container.Invoke(sv.StartFunc)
	if err != nil {
		log.Error("Failed to start:", sv.Name)
		panic(err)
	}
}

func (a *ApplicationContext) provideConfig() {
	for _, c := range a.Configs {
		err := a.Container.Provide(c.Conf)
		if err != nil {
			log.Errorf("Invalid config for %s:%+v with error", c.Prefix, c.Conf, err)
			panic(err)
		}
	}
}

func (a *ApplicationContext) provideServices() {
	for _, s := range a.Services {
		err := a.Container.Provide(s)
		if err != nil {
			panic(err)
		}
	}
}

