package deployer

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/drizzleent/vortex/pkg/client/db"
)

const (
	tableAlgo  = "algorithmstatus"
	client_id  = "client_id"
	vwapColumn = "vwap"
	twapColumn = "twap"
	hftColumn  = "hft"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetActiveList(ctx context.Context) ([]int64, error) {
	builderSelect := sq.Select(client_id).
		From(tableAlgo).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Or{
			sq.Eq{vwapColumn: true},
			sq.Eq{twapColumn: true},
			sq.Eq{hftColumn: true}})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "repository.GetActiveList",
		QueryRaw: query,
	}
	res := make([]int64, 1)
	err = r.db.DB().ScanAllContext(ctx, &res, q, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
