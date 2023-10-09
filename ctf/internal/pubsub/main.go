package pubsub

import (
	"log/slog"

	"github.com/caarlos0/env/v8"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
)

type NatsPubSub struct {
	connection      *nats.Conn
	Url             string `env:"NATS_URL" envDefault:"nats://172.17.0.1:4222"`
	SubscriberTopic string `env:"NATS_SUBSCRIBER_TOPIC" envDefault:"orders"`
	ProducerTopic   string `env:"NATS_PRODUCER_TOPIC" envDefault:"orders"`
}

func NewNats() (*NatsPubSub, error) {
	// Set up a NATS client
	var settings NatsPubSub

	if err := env.Parse(&settings); err != nil {
		return nil, err
	}

	validate := validator.New()

	if err := validate.Struct(settings); err != nil {
		return nil, err
	}

	nats := &settings

	return nats, nil
}

func (n *NatsPubSub) Connect() error {
	conn, err := nats.Connect(n.Url)
	if err != nil {
		return err
	}
	n.connection = conn

	return nil
}

func (n *NatsPubSub) Subscribe(callback func(message []byte) error) error {
	// Subscribe to a topic
	slog.Info("starting handle function")

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	_, err := n.connection.Subscribe(n.SubscriberTopic, func(msg *nats.Msg) {
		slog.Info("got message", "message", msg)
		callback(msg.Data)
		// wg.Done()
	})

	// wg.Wait()
	return err
}

func (n *NatsPubSub) Publish(message []byte) error {
	// Publish a message

	return n.connection.Publish(n.ProducerTopic, message)
}
