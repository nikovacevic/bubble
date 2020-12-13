package user_test

import (
	"testing"

	"github.com/nikovacevic/bubble/pkg/bubble"
	"github.com/nikovacevic/bubble/pkg/user"
)

func TestMockStore(t *testing.T) {
	var store user.Store

	// MockStore should implement the user.Store interface
	// We expect to have 10 users
	store = user.NewMockStore()

	var list []*bubble.User
	var err error

	// Page 1: 5 users
	list, err = store.List(5, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 5 {
		t.Fatalf("expected %d users; got %d", 5, len(list))
	}

	// Page 1: 5 users
	list, err = store.List(5, 2)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 5 {
		t.Fatalf("expected %d users; got %d", 5, len(list))
	}

	// Page 1: 0 users
	list, err = store.List(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 0 {
		t.Fatalf("expected %d users; got %d", 0, len(list))
	}

	// Page 1: 10 users
	list, err = store.List(10, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 10 {
		t.Fatalf("expected %d users; got %d", 10, len(list))
	}

	// Page 1: 10 users
	list, err = store.List(15, 1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(list) != 10 {
		t.Fatalf("expected %d users; got %d", 10, len(list))
	}

	// TODO: Get
	// TODO: Save new (create)
	// TODO: Save existing (update)
}
