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

type OrderStatus string

var (
	Pending   OrderStatus = "pending"
	Completed OrderStatus = "completed"
	Shipped   OrderStatus = "shipped"
	Cancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	Status     OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *Store) StoreOrder(ctx context.Context, order Order) (int, error) {
	var id int
	err := s.db.QueryRow(ctx, "INSERT INTO user_order(user_id, total_price, status) VALUES($1, $2, $3) RETURNING id", order.UserID, order.TotalPrice, string(order.Status)).Scan(&id)
	return id, err
}

func (s *Store) RetrieveOrder(ctx context.Context, id int) (*Order, error) {
	order := new(Order)
	err := s.db.QueryRow(ctx, "SELECT id, user_id, total_price, status, created_at, updated_at FROM user_order WHERE id = $1", id).Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	return order, err
}

func (s *Store) RetrieveOrdersByUserID(ctx context.Context, userID int) ([]*Order, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM user_order WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		order := new(Order)
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Store) UpdateOrder(ctx context.Context, id int, order Order) error {
	_, err := s.db.Exec(ctx, "UPDATE user_order SET status = $2, total_price = $3 WHERE id = $1", id, string(order.Status), order.TotalPrice)
	return err
}

func (s *Store) UpdateOrderStatus(ctx context.Context, id int, status OrderStatus) error {
	_, err := s.db.Exec(ctx, "UPDATE user_order SET status = $2 WHERE id = $1", id, string(status))
	return err
}
