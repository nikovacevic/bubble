package main

import (
	"github.com/nikovacevic/bubble/pkg/message"
	"github.com/nikovacevic/bubble/pkg/user"
	"github.com/nikovacevic/bubble/pkg/venue"
)

type App struct {
	Messages message.Store
	Users    user.Store
	Venues   venue.Store
}

func main() {
	// TODO: connect to postgres

	// TODO: create message.PostgresStore
	// TODO: create user.PostgresStore
	// TODO: create venue.PostgresStore

	// TODO: create App with postgres stores

	// TODO: run App
}
