package tests

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/drizzleent/vortex/internal/api/http/handler"
	"github.com/drizzleent/vortex/internal/model"
	"github.com/drizzleent/vortex/internal/service"
	serviceMocks "github.com/drizzleent/vortex/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestAddClient(t *testing.T) {
	t.Parallel()

	type ApiServiceMock func(mc *minimock.Controller) service.ApiService

	type args struct {
		ctx *gin.Context
	}

	var (
		ctx = &gin.Context{}
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Name()

		serviceErr = fmt.Errorf("service error")
		req        = model.Client{ClientName: name}
		res        = id
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name           string
		args           args
		want           int
		err            error
		apiServiceMock ApiServiceMock
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
			},
			want: 0,
			err:  serviceErr,
			apiServiceMock: func(mc *minimock.Controller) service.ApiService {
				mock := serviceMocks.NewApiServiceMock(mc)
				mock.AddClientMock.Expect(ctx, &req).Return(res, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			apiServiceMock := tt.apiServiceMock(mc)
			api := handler.NewHandler(apiServiceMock)

			api.AddClient(tt.args.ctx)
			require.Equal(t, tt.err, nil)
			require.Equal(t, tt.want, nil)
		})
	}

}
