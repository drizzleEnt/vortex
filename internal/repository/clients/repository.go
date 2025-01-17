package clients

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/drizzleent/vortex/internal/model"
	"github.com/drizzleent/vortex/pkg/client/db"
)

const (
	id         = "id"
	client_id  = "client_id"
	tableCl    = "clients"
	tableAlgo  = "algorithmstatus"
	name       = "client_name"
	createdAt  = "created_at"
	updatedAt  = "updated_at"
	spawnedAt  = "spawned_at"
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

func (r *repo) AddClient(ctx context.Context, client *model.Client) (int64, error) {
	// Add client to clients table
	builderInsert := sq.Insert(tableCl).
		PlaceholderFormat(sq.Dollar).
		Columns(name).
		Values(client.ClientName).
		Suffix("RETURNING id")
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "repository.AddClient to clients",
		QueryRaw: query,
	}
	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)

	if err != nil {
		return 0, err
	}

	// Add client to algorithim table
	builderInsert = sq.Insert(tableAlgo).
		PlaceholderFormat(sq.Dollar).
		Columns(client_id).
		Values(userID).
		Suffix("RETURNING id")
	query, args, err = builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q = db.Query{
		Name:     "repository.AddClient to clients",
		QueryRaw: query,
	}
	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) DeleteClient(ctx context.Context, deleteID int64) error {
	builderInsert := sq.Delete(tableCl).
		Where(sq.Eq{id: deleteID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "repository.DeleteClient",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v tag: %v", err, res)
	}
	return nil
}

func (r *repo) UpdateClient(ctx context.Context, ID int64, info *model.Client) error {
	builderUpdate := sq.Update(tableCl).
		PlaceholderFormat(sq.Dollar).
		Set(name, info.ClientName).
		Set(updatedAt, time.Now()).
		Where(sq.Eq{id: ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "repository.UpdateAlgorithms",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update algorithms: %v tag: %v", err, res)
	}

	return nil
}

func (r *repo) UpdateAlgorithmStatus(ctx context.Context, ID int64, info *model.Algorithms) error {
	builderUpdate := sq.Update(tableAlgo).
		PlaceholderFormat(sq.Dollar).
		Set(twapColumn, info.Twap).
		Set(vwapColumn, info.Vwap).
		Set(hftColumn, info.Hft).
		Where(sq.Eq{client_id: ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "repository.UpdateAlgorithms",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update algorithms: %v tag: %v", err, res)
	}

	return nil
}
