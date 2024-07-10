package app

import (
	"context"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/drizzleent/vortex/internal/api/http/handler"
)

type serviceProvider struct {
	handler api.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler()
	}
	return s.handler
}
