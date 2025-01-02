package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection(config *Config) (PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Db, config.SSLMode,
	)
	log.Printf("postgres connection string is %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Person{}); err != nil {
		return nil, err
	}
	return &PostgresGormDB{db}, nil
}

type PostgresGormDB struct {
	db *gorm.DB
}

func (db *PostgresGormDB) Close() {
	log.Println("closing the postgres connection")
}

func (db *PostgresGormDB) Create(person *Person) (*Person, error) {
	err := db.db.Create(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}
