package products

import (
	"backend-web-commision-kana/internal/repo"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

type Service interface {
	CreateProduct(ctx context.Context, req ReqAdminProducts, userID int32) (*ResAdminProduct, error)
	DeleteProduct(ctx context.Context, productID int32) (*ResAdminProduct, error)
	GetProductAdmin(ctx context.Context, userID int32) ([]ResAdminGetProduct, error)
}

func NewService(repos *repo.Queries, dbs *pgxpool.Pool) Service {
	return &svc{
		repo: repos,
		db:   dbs,
	}
}

func (s *svc) CreateProduct(ctx context.Context, req ReqAdminProducts, userID int32) (*ResAdminProduct, error) {
	err := s.repo.CreateProducts(ctx, repo.CreateProductsParams{
		NamaProducts: req.ProductName,
		Price:        req.ProductPrice,
		UserID:       userID,
	})
	if err != nil {
		return nil, err
	}

	return &ResAdminProduct{
		Msg: "Produk berhasil dibuat!",
	}, nil
}

func (s *svc) DeleteProduct(ctx context.Context, productID int32) (*ResAdminProduct, error) {
	result, err := s.repo.DeleteProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, ErrProductNotFound
	}

	return &ResAdminProduct{
		Msg: "Produk berhasil dihapus!",
	}, nil
}

func (s *svc) GetProductAdmin(ctx context.Context, userID int32) ([]ResAdminGetProduct, error) {
	rows, err := s.repo.GetProductAdmin(ctx, userID)
	if err != nil {
		return nil, err
	}

	productItems := make([]ResAdminGetProduct, 0, len(rows))
	for _, row := range rows {
		productItems = append(productItems, ResAdminGetProduct{
			ProductID:    row.ID,
			ProductName:  row.NamaProducts,
			ProductPrice: row.Price,
		})
	}

	return productItems, nil
}
