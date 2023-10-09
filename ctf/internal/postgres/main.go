package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v8"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Order struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Amount  float32   `json:"amount"`
}

type Postgres struct {
	db       *sql.DB
	Host     string `env:"POSTGRES_HOST" envDefault:"127.0.0.1"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Database string `env:"POSTGRES_DB" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" validate:"required"`
}

func NewPostgres() (*Postgres, error) {
	var settings Postgres

	if err := env.Parse(&settings); err != nil {
		return nil, err
	}

	validate := validator.New()

	if err := validate.Struct(settings); err != nil {
		return nil, err
	}

	pg := &settings

	return pg, nil
}

func (p *Postgres) Connect() error {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			p.Host,
			p.Port,
			p.User,
			p.Password,
			p.Database,
		),
	)

	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) CreateTable() error {
	// Create table
	statement := "CREATE TABLE orders (id uuid PRIMARY KEY, name VARCHAR NOT NULL, address VARCHAR NOT NULL, amount float);"

	_, err := p.db.Exec(statement)
	return err
}

func (p *Postgres) Handle(message []byte) error {
	// Insert a message

	var order Order
	if err := json.Unmarshal(message, &order); err != nil {
		slog.Error("error on creating object out of message", "message", message, "uuid", uuid.New().String())
		return err
	}

	// Prepare the statement and execute it
	statement := fmt.Sprintf(
		"INSERT INTO orders (id, name, address, amount) VALUES ('%s', '%s', '%s', %f);",
		order.ID,
		order.Name,
		order.Address,
		order.Amount,
	)
	_, err := p.db.Exec(statement)

	slog.Info("Attempting to execute SQL statement", "statement", statement)

	// Check if the SQL insert has been performed successfully
	if err != nil {
		slog.Error("could not insert order in database", "order", order)
		return err
	}

	// Log success and exit
	slog.Info("Successfully executed SQL statement", "statement", statement)
	return nil
}
