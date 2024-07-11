package app

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/drizzleent/vortex/internal/config"
	"github.com/drizzleent/vortex/internal/deployer"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	deployer        *deployer.Deployer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebs(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runDeployer()
		if err != nil {
			log.Fatalf("failed to run Deployer: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initDeployer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initDeployer(ctx context.Context) error {
	a.deployer = a.serviceProvider.Deployer(ctx)
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	srv := &http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        a.serviceProvider.Handler(ctx).InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	a.httpServer = srv
	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("server run on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) runDeployer() error {
	a.deployer.Run()
	return nil
}
