package app

import (
	"context"
	"log"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/drizzleent/vortex/internal/api/http/handler"
	"github.com/drizzleent/vortex/internal/config"
	"github.com/drizzleent/vortex/internal/config/env"
	"github.com/drizzleent/vortex/internal/deployer"
	"github.com/drizzleent/vortex/internal/repository"
	clientsRepository "github.com/drizzleent/vortex/internal/repository/clients"
	deployerRepository "github.com/drizzleent/vortex/internal/repository/deployer"
	"github.com/drizzleent/vortex/internal/service"
	clientsService "github.com/drizzleent/vortex/internal/service/clients"
	serviceDeployer "github.com/drizzleent/vortex/internal/service/deployer"
	"github.com/drizzleent/vortex/pkg/client/db"
	"github.com/drizzleent/vortex/pkg/client/db/pg"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	deployerConfig config.DeployerConfig

	dbClient db.Client

	clientRepository   repository.ClientRepository
	deployerRepository repository.DeployerRepository

	service         service.ApiService
	deployerService service.Deployer

	handler api.Handler

	deployer *deployer.Deployer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DeployerConfig() config.DeployerConfig {
	if nil == s.deployerConfig {
		cfg, err := env.NewDeployerConfig()
		if err != nil {
			log.Fatalf("failed to get deployer config: %s", err.Error())
		}
		s.deployerConfig = cfg
	}

	return s.deployerConfig
}

func (s *serviceProvider) PGConfig() config.HTTPConfig {
	if nil == s.pgConfig {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if nil == s.httpConfig {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if nil == s.dbClient {
		cl, err := pg.New(ctx, s.PGConfig().Address())
		if err != nil {
			log.Fatalf("Failed to create db client %s", err.Error())
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("Failed to ping db %s", err.Error())
		}
		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) ApiRepository(ctx context.Context) repository.ClientRepository {
	if nil == s.clientRepository {
		s.clientRepository = clientsRepository.NewRepository(s.DBClient(ctx))
	}
	return s.clientRepository
}

func (s *serviceProvider) ApiService(ctx context.Context) service.ApiService {
	if nil == s.service {
		s.service = clientsService.NewClientService(s.ApiRepository(ctx))
	}
	return s.service
}

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler(s.ApiService(ctx))
	}
	return s.handler
}

func (s *serviceProvider) DeployerRepository(ctx context.Context) repository.DeployerRepository {
	if nil == s.deployerRepository {
		s.deployerRepository = deployerRepository.NewRepository(s.DBClient(ctx))
	}
	return s.deployerRepository
}

func (s *serviceProvider) DeployerService(ctx context.Context) service.Deployer {
	if nil == s.deployerService {
		s.deployerService = serviceDeployer.NewDeployer(s.DeployerRepository(ctx))
	}
	return s.deployerService
}

func (s *serviceProvider) Deployer(ctx context.Context) *deployer.Deployer {
	if nil == s.deployer {
		s.deployer = deployer.NewDeployer(s.DeployerService(ctx), s.DeployerConfig().Duration())
	}
	return s.deployer
}
