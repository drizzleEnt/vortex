package clients

import "context"

func (s *srv) DeleteClient(ctx context.Context, ID int64) error {
	s.r.DeleteClient(ctx, ID)
	return nil
}
