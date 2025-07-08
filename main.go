package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/ponyo877/product-expiry-tracker/adapter"
	"github.com/ponyo877/product-expiry-tracker/openapi"
	"github.com/ponyo877/product-expiry-tracker/repository"
	"github.com/ponyo877/product-expiry-tracker/usecase"
)

func main() {
	// Build DSN from individual environment variables
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("Failed to load timezone:", err)
	}
	c := mysql.Config{
		DBName:    database,
		User:      user,
		Passwd:    password,
		Addr:      fmt.Sprintf("%s:%s", host, port),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	// Connect to database
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	repo := repository.NewRepository(db)
	uc := usecase.NewUseCase(repo)
	h := adapter.NewHandler(uc)
	srv, err := openapi.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	httpPort := ":8080"
	fmt.Printf("Starting server on %s\n", httpPort)

	if err := http.ListenAndServe(httpPort, srv); err != nil {
		log.Fatal(err)
	}
}
