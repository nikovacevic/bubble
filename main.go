package main

import (
	"github.com/nikovacevic/bubble/pkg/account"
	"github.com/nikovacevic/bubble/pkg/message"
	"github.com/nikovacevic/bubble/pkg/venue"
)

type App struct {
	Messages message.Store
	Accounts account.Store
	Venues   venue.Store
}

func main() {
	// TODO: connect to postgres

	// TODO: create message.PostgresStore
	// TODO: create account.PostgresStore
	// TODO: create venue.PostgresStore

	// TODO: create App with postgres stores

	// TODO: run App
}
