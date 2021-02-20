package account_test

import (
	"testing"

	"github.com/nikovacevic/bubble/pkg/account"
	"github.com/nikovacevic/bubble/pkg/bubble"
)

func TestMockStore(t *testing.T) {
	var store account.Store

	// MockStore should implement the account.Store interface
	// We expect to have 10 accounts
	store = account.NewMockStore()

	var list []*bubble.Account
	var err error

	// Page 1: 5 accounts
	list, err = store.List(5, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 5 {
		t.Fatalf("expected %d accounts; got %d", 5, len(list))
	}

	// Page 1: 5 accounts
	list, err = store.List(5, 2)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 5 {
		t.Fatalf("expected %d accounts; got %d", 5, len(list))
	}

	// Page 1: 0 accounts
	list, err = store.List(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 0 {
		t.Fatalf("expected %d accounts; got %d", 0, len(list))
	}

	// Page 1: 10 accounts
	list, err = store.List(10, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 10 {
		t.Fatalf("expected %d accounts; got %d", 10, len(list))
	}

	// Page 1: 10 accounts
	list, err = store.List(15, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 10 {
		t.Fatalf("expected %d accounts; got %d", 10, len(list))
	}

	// TODO: Get
	// TODO: Save new (create)
	// TODO: Save existing (update)
}
