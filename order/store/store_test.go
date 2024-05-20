package store

import (
	"context"
	"testing"

	"github.com/PseudoMera/virtual-store/shared"
)

const (
	projectRootPath = "../../"
)

func TestOrderStore(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	store := NewStore(db.DB())

	var userID int
	err = db.DB().QueryRow(ctx, "INSERT INTO vstore_user(email, password) VALUES($1, $2) RETURNING id", "email", "password").Scan(&userID)
	if err != nil {
		t.Fatal(err)
	}

	order := Order{
		UserID:     userID,
		TotalPrice: 22,
		Status:     Pending,
	}
	id, err := store.StoreOrder(ctx, order)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Fatalf("wanted %d, got %d", 1, id)
	}

	retrievedOrder, err := store.RetrieveOrder(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	if retrievedOrder.UserID != userID {
		t.Fatalf("wanted %d, got %d", userID, retrievedOrder.UserID)
	}
	if retrievedOrder.ID != id {
		t.Fatalf("wanted %d, got %d", id, retrievedOrder.ID)
	}

	orders, err := store.RetrieveOrdersByUserID(ctx, userID)
	if err != nil {
		t.Fatal(err)
	}

	if len(orders) != 1 {
		t.Fatalf("wanted %d, got %d", 1, len(orders))
	}

	if err = store.UpdateOrder(ctx, id, Order{
		Status:     Cancelled,
		TotalPrice: 22,
	}); err != nil {
		t.Fatal(err)
	}

	if err = store.UpdateOrderStatus(ctx, id, Completed); err != nil {
		t.Fatal(err)
	}
}
