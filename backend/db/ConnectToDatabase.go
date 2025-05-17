package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Connection *pgx.Conn

func ConnectToDatabase() {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("No database URL found")
	}

	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database!")

	Connection = conn
}
