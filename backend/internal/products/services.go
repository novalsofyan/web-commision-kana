package products

import (
	"backend-web-commision-kana/internal/repo"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

type Service interface {
}

func NewService(repos *repo.Queries, dbs *pgxpool.Pool) Service {
	return &svc{
		repo: repos,
		db:   dbs,
	}
}

func (s *svc) CreateProduct(ctx context.Context, req ReqAdminProducts) (*ResAdminProduct, error) {
	return nil, nil
}
