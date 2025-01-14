package server

import (
	"log"
	"net/http"

	"github.com/abhishekghoshh/datastore/pkg/dto"
	"github.com/abhishekghoshh/datastore/pkg/postgres"
	"github.com/labstack/echo"
)

type PostgresApi struct {
	db postgres.PostgresDB
}

func NewPostgresApi(postgresDB postgres.PostgresDB) *PostgresApi {
	return &PostgresApi{postgresDB}
}

func (api *PostgresApi) Init(e *echo.Echo) {
	e.GET("/postgres/persons", api.getAllPersons)
	e.GET("/postgres/person/{personId}", api.getPersonsBy)
	e.POST("/postgres/person", api.addPerson)
	e.PUT("/postgres/person/{personId}", api.updatePerson)
	e.DELETE("/postgres/person/{personId}", api.deletePerson)
}

func (*PostgresApi) getAllPersons(c echo.Context) error {
	data := map[string]any{
		"status":         "UP",
		"notImplemented": true,
	}
	return c.JSON(http.StatusOK, data)
}
func (*PostgresApi) getPersonsBy(c echo.Context) error {
	data := map[string]any{
		"status":         "UP",
		"notImplemented": true,
	}
	return c.JSON(http.StatusOK, data)
}
func (api *PostgresApi) addPerson(c echo.Context) error {
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
func (*PostgresApi) updatePerson(c echo.Context) error {
	data := map[string]any{
		"status":         "UP",
		"notImplemented": true,
	}
	return c.JSON(http.StatusOK, data)
}
func (*PostgresApi) deletePerson(c echo.Context) error {
	data := map[string]any{
		"status":         "UP",
		"notImplemented": true,
	}
	return c.JSON(http.StatusOK, data)
}
