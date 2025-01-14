package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhishekghoshh/datastore/pkg/dto"
	_ "github.com/lib/pq"
)

func NewSqlcConnection(config *Config) (PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Db, config.SSLMode,
	)
	log.Printf("postgres connection string is %s", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresSqlcDB{db}, nil
}

type PostgresSqlcDB struct {
	db *sql.DB
}

func (db *PostgresSqlcDB) Close() {
	log.Println("closing the postgres connection")
	db.db.Close()
}

func (db *PostgresSqlcDB) Create(person *dto.Person) (*dto.Person, error) {
	return nil, nil
}
