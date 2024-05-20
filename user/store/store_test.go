package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/PseudoMera/virtual-store/shared"
)

const (
	testEmail    = "testEmail@test.test"
	testPassword = "testing"
	testName     = "tester"

	projectRootPath = "../../"
)

func TestUserStore(t *testing.T) {
	ctx := context.Background()
	db, _, err := shared.SetupPostgresClient(ctx, projectRootPath)
	if err != nil {
		t.Fatal(err)
	}

	store := NewStore(db.DB())
	user := User{
		Email:    testEmail,
		Password: testPassword,
	}

	id, err := store.StoreUser(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Fatalf("wanted %d, got %d", 1, id)
	}

	nUser := User{
		Email: testEmail,
	}
	_, err = store.StoreUser(ctx, nUser)
	if err == nil {
		t.Fatal(err)
	}

	rUser, err := store.RetrieveUser(ctx, testEmail)
	if err != nil {
		t.Fatal(err)
	}

	if rUser.Email != testEmail {
		t.Fatalf("wanted %s, got %s", testEmail, rUser.Email)
	}

	profile := Profile{
		UserID: id,
		Name:   testName,
	}

	_, err = store.StoreUserProfile(ctx, profile)
	if err != nil {
		t.Fatal(err)
	}

	prof, err := store.RetrieveUserProfile(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	if prof.Name != profile.Name {
		t.Fatalf("wanted %s, got %s", testName, prof.Name)
	}
	if prof.UserID != id {
		fmt.Printf("%v\n", prof)
		t.Fatalf("wanted %d, got %d", id, prof.UserID)
	}

	updatedProfile := Profile{
		UserID: id,
		Name:   testName + "2",
	}
	if err = store.UpdateUserProfile(ctx, updatedProfile); err != nil {
		t.Fatal(err)
	}
}
