package main

import (
	"github.com/nikovacevic/bubble/pkg/account"
	"github.com/nikovacevic/bubble/pkg/destination"
	"github.com/nikovacevic/bubble/pkg/message"
)

type App struct {
	Messages     message.Store
	Accounts     account.Store
	Destinations destination.Store
}

func main() {
	// TODO: connect to postgres

	// TODO: create message.PostgresStore
	// TODO: create account.PostgresStore
	// TODO: create venue.PostgresStore

	// TODO: create App with postgres stores

	// TODO: run App
}
