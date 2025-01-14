package postgres

import "github.com/abhishekghoshh/datastore/pkg/dto"

type PostgresDB interface {
	Create(person *dto.Person) (*dto.Person, error)
	Close()
}
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	Db       string
	SSLMode  string
}
