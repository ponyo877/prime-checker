package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/ponyo877/product-expiry-tracker/db/generated_sql"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func NewDatabaseConnection(config DatabaseConfig) (*sql.DB, error) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	c := mysql.Config{
		DBName:    config.Database,
		User:      config.User,
		Passwd:    config.Password,
		Addr:      fmt.Sprintf("%s:%s", config.Host, config.Port),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func NewQueries(db *sql.DB) *generated_sql.Queries {
	return generated_sql.New(db)
}
