package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/abhishekghoshh/datastore/pkg/db"
	"github.com/abhishekghoshh/datastore/pkg/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoConnection(config *db.Config) (db.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	dsn := fmt.Sprintf(
		"mongodb://%s:%s",
		config.Host, config.Port,
	)
	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		defer cancel()
		return nil, err
	}
	client.Ping(ctx, readpref.Primary())
	return &MongoDb{
		ctx:    ctx,
		cancel: cancel,
		client: client,
		db:     config.Db,
	}, nil
}

type MongoDb struct {
	ctx    context.Context
	cancel context.CancelFunc
	client *mongo.Client
	db     string
}

func (db *MongoDb) Close() {
	db.cancel()
}

func (db *MongoDb) Create(person *dto.Person) (*dto.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	personMongo := db.convertFromPerson(person)
	collection := db.client.Database(db.db).Collection(personMongo.CollectionName())
	lastId, err := collection.InsertOne(ctx, personMongo)
	if err != nil {
		return nil, err
	}
	person = db.converToPerson(personMongo)
	person.ID = lastId.InsertedID.(bson.ObjectID).Hex()
	return person, nil
}

type PersonMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string
	UserName string
	Age      int64
	Email    string
	Notes    []Note
}

func (PersonMongo) CollectionName() string {
	return "people"
}

type Note struct {
	Content string
}

func (MongoDb) convertFromPerson(person *dto.Person) *PersonMongo {
	var notes []Note
	if person.Notes != nil {
		for i := 0; i < len(person.Notes); i++ {
			notes = append(notes, Note{Content: person.Notes[i].Content})
		}
	}

	return &PersonMongo{
		Name:     person.Name,
		UserName: person.UserName,
		Email:    person.Email,
		Age:      person.Age,
		Notes:    notes,
	}
}
func (MongoDb) converToPerson(person *PersonMongo) *dto.Person {
	var notes []dto.Note
	if person.Notes != nil {
		for i := 0; i < len(person.Notes); i++ {
			noteMongo := person.Notes[i]
			notes = append(notes, dto.Note{
				Content: noteMongo.Content,
			})
		}
	}
	return &dto.Person{
		ID:       person.ID.Hex(),
		Name:     person.Name,
		UserName: person.UserName,
		Email:    person.Email,
		Age:      person.Age,
		Notes:    notes,
	}
}
