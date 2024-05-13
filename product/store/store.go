package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Store) StoreProduct(ctx context.Context, product Product) (int, error) {
	var id int
	err := s.db.QueryRow(ctx, "INSERT INTO product(name, price, stock) VALUES($1, $2, $3) RETURNING id", product.Name, product.Price, product.Stock).Scan(&id)
	return id, err
}

func (s *Store) RetrieveProduct(ctx context.Context, id int) (*Product, error) {
	product := new(Product)
	err := s.db.QueryRow(ctx, "SELECT * FROM product WHERE id = $1", id).Scan(&product)
	return product, err
}

func (s *Store) RetrieveProducts(ctx context.Context, name string) ([]*Product, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM product WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		product := new(Product)
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *Store) UpdateProduct(ctx context.Context, product Product) error {
	_, err := s.db.Exec(ctx, "UPDATE product SET name = $2, price = $3, stock = $4 WHERE id = $1", product.ID, product.Name, product.Price, product.Stock)
	return err
}

func (s *Store) UpdateProductStock(ctx context.Context, id, stock int) error {
	_, err := s.db.Exec(ctx, "UPDATE product SET stock = $2 WHERE id = $1", id, stock)
	return err
}
