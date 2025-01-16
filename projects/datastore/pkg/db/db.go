package db

import "github.com/abhishekghoshh/datastore/pkg/dto"

type DB interface {
	Create(person *dto.Person) (*dto.Person, error)
	Close()
}
type Config struct {
	Host     string
	Port     string
	Db       string
	User     string
	Password string
	SSLMode  string
}
