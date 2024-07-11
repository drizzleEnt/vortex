package deployer

import (
	"sync"

	"github.com/drizzleent/vortex/internal/repository"
)

type deployer struct {
	m    sync.RWMutex
	pods map[string]struct{}
	repo repository.DeployerRepository
}

func NewDeployer(r repository.DeployerRepository) *deployer {
	return &deployer{
		pods: make(map[string]struct{}),
		repo: r,
	}
}
