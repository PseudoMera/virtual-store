package grpc

import (
	context "context"
	"errors"

	"github.com/PseudoMera/virtual-store/product/store"
)

var (
	errEmptyName  = errors.New("name field cannot be empty")
	errEmptyPrice = errors.New("price field cannot be empty")
	errEmptyStock = errors.New("stock field cannot be empty")
	errEmptyID    = errors.New("id field cannot be empty")
)

type ProductServer struct {
	db *store.Store
	UnimplementedProductServiceServer
}

func NewProductServer(db *store.Store) *ProductServer {
	return &ProductServer{
		db: db,
	}
}

func (ps *ProductServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	if req.Name == "" {
		return nil, errEmptyName
	}
	if req.Price == 0.0 {
		return nil, errEmptyPrice
	}
	if req.Stock == 0 {
		return nil, errEmptyStock
	}

	id, err := ps.db.StoreProduct(ctx, store.Product{
		Name:  req.Name,
		Price: float64(req.Price),
		Stock: int(req.Stock),
	})
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		Id: int64(id),
	}, nil
}

func (ps *ProductServer) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	if req.Id == 0 {
		return nil, errEmptyID
	}

	product, err := ps.db.RetrieveProduct(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &GetProductResponse{
		Product: &Product{
			Id:    int64(product.ID),
			Name:  product.Name,
			Stock: int32(product.Stock),
			Price: float32(product.Price),
		},
	}, nil
}

func (ps *ProductServer) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	if req.Name == "" {
		return nil, errEmptyName
	}

	products, err := ps.db.RetrieveProducts(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	parsedProducts := make([]*Product, len(products))
	for i := range products {
		parsedProducts = append(parsedProducts, &Product{
			Id:    int64(products[i].ID),
			Name:  products[i].Name,
			Price: float32(products[i].Price),
			Stock: int32(products[i].Stock),
		})
	}

	return &GetProductsResponse{
		Products: parsedProducts,
	}, nil
}

func (ps *ProductServer) UpdateProductRequest(ctx context.Context, req *Product) (*SuccessResponse, error) {
	id, name, price, stock := req.Id, req.Name, req.Price, req.Stock
	if name == "" {
		return nil, errEmptyName
	}
	if price == 0.0 {
		return nil, errEmptyPrice
	}
	if stock == 0 {
		return nil, errEmptyStock
	}

	err := ps.db.UpdateProduct(ctx, store.Product{
		ID:    int(id),
		Name:  name,
		Price: float64(price),
		Stock: int(stock),
	})
	if err != nil {
		return nil, err
	}

	return &SuccessResponse{
		Msg: "Success!",
	}, nil
}

func (ps *ProductServer) UpdateProductStock(ctx context.Context, req *UpdateProductStockRequest) (*SuccessResponse, error) {
	id, stock := req.Id, req.Stock
	if stock == 0 {
		return nil, errEmptyStock
	}

	err := ps.db.UpdateProductStock(ctx, int(id), int(stock))
	if err != nil {
		return nil, err
	}

	return &SuccessResponse{
		Msg: "Success!",
	}, nil
}
