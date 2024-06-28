package grpc

import (
	context "context"

	"github.com/PseudoMera/virtual-store/order/store"
)

type OrderServer struct {
	db *store.Store
	UnimplementedOrderServiceServer
}

// NewOrderServer returns a GRPC server with the given database
func NewOrderServer(db *store.Store) *OrderServer {
	return &OrderServer{
		db: db,
	}
}

func (os *OrderServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	if req.UserID == 0 {
		return nil, nil
	}
	if req.Status == "" {
		return nil, nil
	}
	if req.TotalPrice == 0.0 {
		return nil, nil
	}

	id, err := os.db.StoreOrder(ctx, store.Order{
		UserID:     int(req.UserID),
		Status:     store.OrderStatus(req.Status),
		TotalPrice: float64(req.TotalPrice),
	})

	return &CreateOrderResponse{
		Id: int64(id),
	}, err
}

func (os *OrderServer) GetOrder(ctx context.Context, req *GetOrderRequest) (*Order, error) {
	if req.Id == 0 {
		return nil, nil
	}

	order, err := os.db.RetrieveOrder(ctx, int(req.Id))
	return &Order{
		Id:         int64(order.ID),
		UserID:     int64(order.UserID),
		TotalPrice: float32(order.TotalPrice),
		Status:     string(order.Status),
	}, err
}

func (os *OrderServer) GetOrdersByUser(ctx context.Context, req *GetOrdersByUserRequest) (*GetOrdersByUserResponse, error) {
	if req.UserID == 0 {
		return nil, nil
	}

	orders, err := os.db.RetrieveOrdersByUserID(ctx, int(req.UserID))
	if err != nil {
		return nil, err
	}

	parsedOrders := make([]*Order, len(orders))
	for i := range orders {
		parsedOrders[i] = &Order{
			Id:         int64(orders[i].ID),
			UserID:     int64(orders[i].UserID),
			TotalPrice: float32(orders[i].TotalPrice),
			Status:     string(orders[i].Status),
		}
	}

	return &GetOrdersByUserResponse{
		Orders: parsedOrders,
	}, nil
}

func (os *OrderServer) UpdateOrder(ctx context.Context, req *UpdateOrderRequest) (*SuccessResponse, error) {
	if req.Id == 0 {
		return nil, nil
	}
	if req.Status == "" {
		return nil, nil
	}
	if req.TotalPrice == 0.0 {
		return nil, nil
	}

	if err := os.db.UpdateOrder(ctx, int(req.Id), store.Order{
		ID:         int(req.Id),
		TotalPrice: float64(req.TotalPrice),
		Status:     store.OrderStatus(req.Status),
	}); err != nil {
		return nil, err
	}

	return &SuccessResponse{
		Msg: "Success!",
	}, nil
}

func (os *OrderServer) UpdateOrderStatus(ctx context.Context, req *UpdateOrderStatusRequest) (*SuccessResponse, error) {
	if req.Id == 0 {
		return nil, nil
	}
	if req.Status == "" {
		return nil, nil
	}
	if err := os.db.UpdateOrderStatus(ctx, int(req.Id), store.OrderStatus(req.Status)); err != nil {
		return nil, err
	}

	return &SuccessResponse{
		Msg: "Success!",
	}, nil
}
