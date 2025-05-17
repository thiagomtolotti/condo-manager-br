package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDatabase() *pgx.Conn {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("No database URL found")
	}

	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database!")
	return conn
}
