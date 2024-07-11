package clients

import (
	"context"

	"github.com/drizzleent/vortex/internal/model"
)

func (s *srv) UpdateAlgorithmStatus(ctx context.Context, ID int64, algos *model.Algorithms) error {
	err := s.r.UpdateAlgorithmStatus(ctx, ID, algos)
	if err != nil {
		return err
	}
	return nil
}
