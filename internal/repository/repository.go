package repository

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

type ClientRepository interface {
	AddClient(context.Context, *model.Client) (int64, error)
	DeleteClient(context.Context, int64) error
	UpdateClient(context.Context, int64, *model.Client) error
	UpdateAlgorithmStatus(context.Context, int64, *model.Algorithms) error
}

type DeployerRepository interface {
	GetActiveList(context.Context) ([]int64, []int64, error)
}
