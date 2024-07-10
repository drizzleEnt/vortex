package clients

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

func (s *srv) AddClient(ctx context.Context, client *model.Client) (int64, error) {
	resp, err := s.r.AddClient(ctx, client)
	if err != nil {
		return 0, nil
	}

	return resp, nil
}
