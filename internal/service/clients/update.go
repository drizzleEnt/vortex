package clients

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

func (s *srv) UpdateClient(ctx context.Context, ID int64, info *model.Client) error {
	err := s.r.UpdateClient(ctx, ID, info)
	if err != nil {
		return err
	}

	return nil
}
