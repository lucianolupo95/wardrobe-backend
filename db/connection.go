package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("No pude cargar el .env en la carpeta actual")
	}

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, dbname, sslmode,
    )

    conn, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    fmt.Println("Database connected successfully!")
    DB = conn
}
