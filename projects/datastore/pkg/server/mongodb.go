package server

import (
	"log"
	"net/http"

	"github.com/abhishekghoshh/datastore/pkg/db"
	"github.com/abhishekghoshh/datastore/pkg/dto"
	"github.com/labstack/echo"
)

type MongoDBApi struct {
	db db.DB
}

func NewMongoDbApi(db db.DB) Server {
	return &MongoDBApi{db}
}

func (api *MongoDBApi) Init(e *echo.Echo) {
	e.GET("/mongo/persons", api.getAllPersons)
	e.GET("/mongo/person/{personId}", api.getPersonsBy)
	e.POST("/mongo/person", api.addPerson)
	e.PUT("/mongo/person/{personId}", api.updatePerson)
	e.DELETE("/mongo/person/{personId}", api.deletePerson)
}

func (*MongoDBApi) getAllPersons(c echo.Context) error {
	return nil
}
func (*MongoDBApi) getPersonsBy(c echo.Context) error {
	return nil
}
func (api *MongoDBApi) addPerson(c echo.Context) error {
	var person dto.Person
	if err := c.Bind(&person); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]any{"stutus": "DataNotInserted", "error": err.Error()})
	}
	p, err := api.db.Create(&person)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]any{"stutus": "DataNotInserted", "error": err.Error()})
	}
	data := map[string]any{
		"status": "DataInserted",
		"additionalData": map[string]any{
			"userId":   p.ID,
			"userName": p.UserName,
		},
	}
	return c.JSON(http.StatusOK, data)
}
func (*MongoDBApi) updatePerson(c echo.Context) error {
	return nil
}
func (*MongoDBApi) deletePerson(c echo.Context) error {
	return nil
}
