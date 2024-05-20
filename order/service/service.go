package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/PseudoMera/virtual-store/order/store"
)

var (
	errEmptyUserID     = errors.New("user id field cannot be empty")
	errEmptyTotalPrice = errors.New("total price field cannot be empty")
	errEmptyId         = errors.New("id field cannot be empty")
	errEmptyStatus     = errors.New("status field cannot be empty")
)

type OrderService struct {
	db     *store.Store
	logger *slog.Logger
}

func NewOrderService(db *store.Store, logger *slog.Logger) *OrderService {
	return &OrderService{
		db:     db,
		logger: logger,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, userID int, totalPrice float64, status store.OrderStatus) (int, error) {
	if userID == 0 {
		o.logger.Info("error at CreateOrder", slog.String("error", errEmptyId.Error()))
		return 0, errEmptyUserID
	}
	if totalPrice == 0.0 {
		o.logger.Info("error at CreateOrder", slog.String("error", errEmptyTotalPrice.Error()))
		return 0, errEmptyTotalPrice
	}
	if status == "" {
		o.logger.Info("error at CreateOrder", slog.String("error", errEmptyStatus.Error()))
		return 0, errEmptyStatus
	}

	return o.db.StoreOrder(ctx, store.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     status,
	})
}

func (o *OrderService) GetOrder(ctx context.Context, id int) (*store.Order, error) {
	if id == 0 {
		o.logger.Info("error at GetOrder", slog.String("error", errEmptyId.Error()))
		return nil, errEmptyId
	}

	return o.db.RetrieveOrder(ctx, id)
}

func (o *OrderService) GetOrdersByUser(ctx context.Context, userID int) ([]*store.Order, error) {
	if userID == 0 {
		o.logger.Info("error at GetOrdersByUser", slog.String("error", errEmptyUserID.Error()))
		return nil, errEmptyUserID
	}

	return o.db.RetrieveOrdersByUserID(ctx, userID)
}

func (o *OrderService) UpdateOrder(ctx context.Context, id int, status store.OrderStatus, totalPrice float64) error {
	if id == 0 {
		o.logger.Info("error at UpdateOrder", slog.String("error", errEmptyId.Error()))
		return errEmptyId
	}
	if status == "" {
		o.logger.Info("error at UpdateOrder", slog.String("error", errEmptyStatus.Error()))
		return errEmptyStatus
	}
	if totalPrice == 0.0 {
		o.logger.Info("error at UpdateOrder", slog.String("error", errEmptyTotalPrice.Error()))
		return errEmptyTotalPrice
	}

	return o.db.UpdateOrder(ctx, id, store.Order{
		ID:         id,
		Status:     status,
		TotalPrice: totalPrice,
	})
}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, id int, status store.OrderStatus) error {
	if id == 0 {
		o.logger.Info("error at UpdateOrder", slog.String("error", errEmptyId.Error()))
		return errEmptyId
	}
	if status == "" {
		o.logger.Info("error at UpdateOrder", slog.String("error", errEmptyStatus.Error()))
		return errEmptyStatus
	}

	return o.db.UpdateOrderStatus(ctx, id, status)
}
