package account_test

import (
	"fmt"
	"testing"

	"github.com/nikovacevic/bubble/pkg/account"
	"github.com/nikovacevic/bubble/pkg/bubble"
)

func TestMockStore(t *testing.T) {
	var store account.Store

	// MockStore should implement the account.Store interface
	// We expect to have 10 accounts
	store = account.NewMockStore().Seed()

	var list []*bubble.Account
	var err error

	list, err = store.ListAccounts()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	for _, acct := range list {
		fmt.Println(acct)
	}
	if len(list) != 10 {
		t.Fatalf("expected %d accounts; got %d", 10, len(list))
	}
}
