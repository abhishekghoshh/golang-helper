package main

import (
	"net/http"

	"github.com/abhishekghoshh/datastore/pkg/config"
	"github.com/abhishekghoshh/datastore/pkg/db"
	"github.com/abhishekghoshh/datastore/pkg/mongodb"
	"github.com/abhishekghoshh/datastore/pkg/postgres"
	"github.com/abhishekghoshh/datastore/pkg/server"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	// creating an instance of the config
	cfg := config.New()

	// setting up the postgress connection and endpoints
	var postgresCfg db.Config
	cfg.Decode("datastore.postgres", &postgresCfg)
	// opening the postgres connection
	postgresDB, err := postgres.NewGormConnection(&postgresCfg)
	if err != nil {
		panic(err)
	}
	// closing the connection and init all the apis
	defer postgresDB.Close()
	server.NewPostgresApi(postgresDB).Init(e)

	// setting up the mongodb connection and endpoints
	var mongoCfg db.Config
	cfg.Decode("datastore.mongo", &mongoCfg)
	// opening the mongo connection
	mongoDb, err := mongodb.NewMongoConnection(&mongoCfg)
	if err != nil {
		panic(err)
	}
	// closing the connection and init all the apis
	defer mongoDb.Close()
	server.NewMongoDbApi(mongoDb).Init(e)

	// spinning up the server
	e.GET("/health", HealthCheck)
	e.Logger.SetLevel(log.INFO)
	e.Logger.Fatal(e.Start(":" + cfg.GetString("server.port")))

}
func HealthCheck(c echo.Context) error {
	type HealthStatus struct {
		Status string            `json:"status"`
		Errors map[string]string `json:"errors,omitempty"`
	}
	health := HealthStatus{
		Status: "UP",
		Errors: make(map[string]string),
	}
	return c.JSON(http.StatusOK, health)
}
