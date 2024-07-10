package service

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

type ApiService interface {
	AddClient(context.Context, *model.Client) (int64, error)
	DeleteClient(context.Context, int64) error
	UpdateClient(context.Context)
	UpdateAlgorithmStatus(context.Context)
}
