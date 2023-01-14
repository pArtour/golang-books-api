package model

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Repository struct {
	BooksRepository *booksRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BooksRepository: NewBookRepository(db),
	}
}

func Init() (*Repository, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	cfgInfo := Connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	db, err := sql.Open("postgres", configToString(&cfgInfo))

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}
	log.Print("Database connected")
	fmt.Printf("config: %s", configToString(&cfgInfo))
	err = db.Ping()

	if err != nil {
		log.Fatalf("pinging: %v", err)
		return nil, err
	}
	log.Print("Database pinged")

	repository := NewRepository(db)
	return repository, nil

}

func configToString(cfg *Connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
}
