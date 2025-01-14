package postgres

import (
	"fmt"
	"log"

	"github.com/abhishekghoshh/datastore/pkg/dto"
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
	if err := db.AutoMigrate(&PersonGorm{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&NoteGorm{}); err != nil {
		return nil, err
	}
	return &PostgresGormDB{db}, nil
}

type PersonGorm struct {
	gorm.Model
	Name     string
	UserName string `gorm:"unique;not null;type:varchar(25);default:null"`
	Age      int64
	Email    string     `gorm:"unique;not null;type:varchar(100);default:null"`
	Notes    []NoteGorm `gorm:"foreignKey:PersonID"`
}

func (PersonGorm) TableName() string {
	return "people"
}

type NoteGorm struct {
	gorm.Model
	Content  string `gorm:"not null;type:varchar(1000);default:null"`
	PersonID uint   `gorm:"not null"` // Foreign key to link to the Person
}

func (NoteGorm) TableName() string {
	return "notes"
}

type PostgresGormDB struct {
	db *gorm.DB
}

func (db *PostgresGormDB) Close() {
	log.Println("closing the postgres connection")
}

func (db *PostgresGormDB) Create(person *dto.Person) (*dto.Person, error) {
	personGorm := db.convertFromPerson(person)
	err := db.db.Create(personGorm).Error
	if err != nil {
		return nil, err
	}
	return db.converToPerson(personGorm), nil
}

func (PostgresGormDB) convertFromPerson(person *dto.Person) *PersonGorm {
	var notes []NoteGorm
	if person.Notes != nil {
		for i := 0; i < len(person.Notes); i++ {
			notes = append(notes, NoteGorm{Content: person.Notes[i].Content})
		}
	}
	return &PersonGorm{
		Name:     person.Name,
		UserName: person.UserName,
		Email:    person.Email,
		Age:      person.Age,
		Notes:    notes,
	}
}
func (PostgresGormDB) converToPerson(person *PersonGorm) *dto.Person {
	var notes []dto.Note
	if person.Notes != nil {
		for i := 0; i < len(person.Notes); i++ {
			noteGorm := person.Notes[i]
			notes = append(notes, dto.Note{
				NoteId:  noteGorm.ID,
				Content: noteGorm.Content,
			})
		}
	}
	return &dto.Person{
		ID:       person.ID,
		Name:     person.Name,
		UserName: person.UserName,
		Email:    person.Email,
		Age:      person.Age,
		Notes:    notes,
	}
}
