package store

import (
	"context"
	"testing"

	"github.com/PseudoMera/virtual-store/shared"
)

const (
	projectRootPath = "../../"
)

func TestProductStore(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	store := NewStore(db.DB())

	product := Product{
		Name:  "product",
		Price: 22.5,
		Stock: 2,
	}
	id, err := store.StoreProduct(ctx, product)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Fatalf("wanted %d, got %d", 1, id)
	}

	retrievedProduct, err := store.RetrieveProduct(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	if retrievedProduct.ID != id {
		t.Fatalf("wanted %d, got %d", id, retrievedProduct.ID)
	}

	products, err := store.RetrieveProducts(ctx, "product")
	if err != nil {
		t.Fatal(err)
	}

	if products[0].ID != id {
		t.Fatalf("wanted %d, got %d", id, products[0].ID)
	}

	if err = store.UpdateProduct(ctx, Product{
		ID:    id,
		Name:  "product2",
		Stock: 3,
		Price: 120,
	}); err != nil {
		t.Fatal(err)
	}

	if err = store.UpdateProductStock(ctx, id, 22); err != nil {
		t.Fatal(err)
	}
}
