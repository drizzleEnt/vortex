package clients

import "github.com/drizzleent/vortex/internal/repository"

type srv struct {
	r repository.ClientRepository
}

func NewClientService(repo repository.ClientRepository) *srv {
	return &srv{
		r: repo,
	}
}
