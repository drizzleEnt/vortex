package service

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

type ApiService interface {
	AddClient(context.Context, *model.Client) (int64, error)
	DeleteClient(context.Context, int64) error
	UpdateClient(context.Context, int64, *model.Client) error
	UpdateAlgorithmStatus(context.Context, int64, *model.Algorithms) error
}

type Deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
	GetActiveList(context.Context) ([]int64, []int64, error)
}
