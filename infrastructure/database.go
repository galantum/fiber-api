package infrastructure

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	var err error
	DB, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/fiber-api")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to PostgreSQL")
}
