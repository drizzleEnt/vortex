package clients

import "github.com/drizzleent/vortex/pkg/client/db"

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) AddClient() {

}

func (r *repo) DeleteClient() {

}

func (r *repo) UpdateClient() {

}

func (r *repo) UpdateAlgorithmStatus() {

}
