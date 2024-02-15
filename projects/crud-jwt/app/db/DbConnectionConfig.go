package db

import (
	"crud/app/config"
	"crud/app/domain"
	lgr "crud/app/logger"

	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbConnection() *gorm.DB {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
			ParameterizedQueries:      false,
		},
	)
	dbUser := config.Get("db.connection.user")
	dbPasswd := config.Get("db.connection.password")
	dbAddr := config.Get("db.connection.address")
	dbPort := config.Get("db.connection.port")
	dbName := config.Get("db.connection.database")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	lgr.Info("datasource is " + dataSource)
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		lgr.Error("Error connecting to database : error : " + err.Error())
		panic(err)
	}
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Account{})
	db.AutoMigrate(&domain.LoggedInUser{})
	return db
}
