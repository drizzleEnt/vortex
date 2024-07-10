package clients

import "context"

func (s *srv) UpdateClient(ctx context.Context) {
	s.r.UpdateClient(ctx)
}
