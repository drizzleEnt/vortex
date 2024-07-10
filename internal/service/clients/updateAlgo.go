package clients

import "context"

func (s *srv) UpdateAlgorithmStatus(ctx context.Context) {
	s.r.UpdateAlgorithmStatus(ctx)
}
