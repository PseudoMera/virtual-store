package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/PseudoMera/virtual-store/product/store"
)

var (
	errEmptyName  = errors.New("name field cannot be empty")
	errEmptyPrice = errors.New("price field cannot be empty")
	errEmptyStock = errors.New("stock field cannot be empty")
	errEmptyID    = errors.New("id field cannot be empty")
)

type ProductService struct {
	db     *store.Store
	logger *slog.Logger
}

// NewProductService returns a ProductService with the given db and logger.
// The product service can be used to interact with the database through
// the Store struct. The service also has validation for the methods
// so it's better to use this than directly interacting with the Store struct.
func NewProductService(db *store.Store, logger *slog.Logger) *ProductService {
	return &ProductService{
		db:     db,
		logger: logger,
	}
}

// CreateProduct stores a new product with the given name, price and stock.
func (p *ProductService) CreateProduct(ctx context.Context, name string, price float64, stock int) (int, error) {
	if name == "" {
		p.logger.Info("error at CreateProduct", slog.String("error", errEmptyName.Error()))
		return 0, errEmptyName
	}
	if price == 0.0 {
		p.logger.Info("error at CreateProduct", slog.String("error", errEmptyPrice.Error()))
		return 0, errEmptyPrice
	}
	if stock == 0 {
		p.logger.Info("error at CreateProduct", slog.String("error", errEmptyStock.Error()))
		return 0, errEmptyStock
	}

	return p.db.StoreProduct(ctx, store.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	})
}

// GetProduct returns the product associated with the given id if it exists.
func (p *ProductService) GetProduct(ctx context.Context, id int) (*store.Product, error) {
	if id == 0 {
		p.logger.Info("error at CreateProduct", slog.String("error", errEmptyID.Error()))
		return nil, errEmptyID
	}

	return p.db.RetrieveProduct(ctx, id)
}

// GetProducts returns a slice of products associated with the given name.
func (p *ProductService) GetProducts(ctx context.Context, name string) ([]*store.Product, error) {
	if name == "" {
		p.logger.Info("error at GetProducts", slog.String("error", errEmptyName.Error()))
		return nil, errEmptyName
	}

	return p.db.RetrieveProducts(ctx, name)
}

// UpdateProduct updates a product with the given id.
func (p *ProductService) UpdateProduct(ctx context.Context, id int, name string, price float64, stock int) error {
	if name == "" {
		p.logger.Info("error at UpdateProduct", slog.String("error", errEmptyName.Error()))
		return errEmptyName
	}
	if price == 0.0 {
		p.logger.Info("error at UpdateProduct", slog.String("error", errEmptyPrice.Error()))
		return errEmptyPrice
	}
	if stock == 0 {
		p.logger.Info("error at UpdateProduct", slog.String("error", errEmptyStock.Error()))
		return errEmptyStock
	}
	if id == 0 {
		p.logger.Info("error at UpdateProduct", slog.String("error", errEmptyID.Error()))
		return errEmptyID
	}

	return p.db.UpdateProduct(ctx, store.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
}

// UpdateProductStock updates the stock of the product associated with the specified id.
func (p *ProductService) UpdateProductStock(ctx context.Context, id int, stock int) error {
	if stock == 0 {
		p.logger.Info("error at UpdateProductStock", slog.String("error", errEmptyStock.Error()))
		return errEmptyStock
	}
	if id == 0 {
		p.logger.Info("error at UpdateProductStock", slog.String("error", errEmptyID.Error()))
		return errEmptyID
	}

	return p.db.UpdateProductStock(ctx, id, stock)
}
