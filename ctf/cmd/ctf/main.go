package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/rutgerhartog/lab/ctf/internal/postgres"
	"github.com/rutgerhartog/lab/ctf/internal/pubsub"
)

func main() {

	// Set up a NATS subscription
	nats, err := pubsub.NewNats()

	if err != nil {
		slog.Error("could not setup NATS")
		os.Exit(1)
	}

	if err := nats.Connect(); err != nil {
		slog.Error("could not connect to NATS server")
		os.Exit(1)
	}

	slog.Info("Set up connection to NATS", "nats", nats)

	// Set up a PGSQL connection
	pg, err := postgres.NewPostgres()

	if err != nil {
		slog.Error("could not parse the database settings")
		os.Exit(1)
	}

	if err := pg.Connect(); err != nil {
		slog.Error("could not set up a connection to the database", "error", err)
		os.Exit(1)
	}

	if err := pg.CreateTable(); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			slog.Error("could not create order table", "error", err)
			os.Exit(1)
		}
	}

	/*
		Finally, set up a listener to NATS that calls a handler function whenever it receives a message.
		Specifically, this function tries to create an order object from the NATS message and tries to insert it into the database.
	*/

	slog.Info("starting app...")
	nats.Subscribe(pg.Handle)
	select {}
}
