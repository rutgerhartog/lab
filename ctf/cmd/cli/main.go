package main

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/rutgerhartog/lab/ctf/internal/postgres"
	"github.com/rutgerhartog/lab/ctf/internal/pubsub"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{Use: "orders"}

	var id string
	var name string
	var address string
	var amount float32

	nats, err := pubsub.NewNats()
	if err != nil {
		slog.Error("could not parse NATS settings", "error", err)
	}

	if err := nats.Connect(); err != nil {
		slog.Error("could not connect to NATS", "error", err)
	}

	slog.Info("Set up connection to NATS", "nats", nats)

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Push an order to the database",
		Run: func(cmd *cobra.Command, args []string) {
			if len(id) == 0 {
				id = uuid.New().String()
			}
			order := postgres.Order{ID: uuid.MustParse(id), Name: name, Address: address, Amount: amount}
			rawOrder, _ := json.Marshal(order)
			if err := nats.Publish(rawOrder); err != nil {
				slog.Error("an error occurred", "error", err)
			} else {
				slog.Info("sent message", "order", order)
			}
		},
	}

	pushCmd.Flags().StringVarP(&id, "id", "i", "", "Order identifier")
	pushCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the client")
	pushCmd.Flags().StringVarP(&address, "address", "d", "", "Address of the client")
	pushCmd.Flags().Float32VarP(&amount, "amount", "m", 0, "Transaction amount")

	rootCmd.AddCommand(pushCmd)

	if err := rootCmd.Execute(); err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

}
