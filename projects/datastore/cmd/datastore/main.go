package main

import (
	"github.com/abhishekghoshh/datastore/pkg/config"
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
	var postgresCfg postgres.Config
	cfg.Decode("datastore.postgres", &postgresCfg)
	// opening the postgres connection
	postgresDB, err := postgres.NewGormConnection(&postgresCfg)
	if err != nil {
		panic(err)
	}
	// closing the connection
	defer postgresDB.Close()
	server.NewPostgresApi(postgresDB).Init(e)

	// spinning up the server
	e.Logger.SetLevel(log.INFO)
	e.Logger.Fatal(e.Start(":" + cfg.GetString("server.port")))
}
