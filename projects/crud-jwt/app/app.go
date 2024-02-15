package app

import (
	"crud/app/config"
	"crud/app/controller"
	"crud/app/db"
	"crud/app/logger"
	"crud/app/middleware"
	"crud/app/service"

	"crud/app/util"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	authMiddleWare := middleware.NewAuthMiddleware()
	dbConnection := db.DbConnection()
	encoder := util.Encoder(config.Get("encoder.secret"))

	accountRepo := db.NewAccountRepo(dbConnection)
	userRepo := db.NewUserRepo(dbConnection)
	tokeRepo := db.NewTokenRepo(dbConnection)

	accountService := service.NewAccountService(accountRepo, userRepo, encoder)
	sessionService := service.NewSessionService(userRepo, tokeRepo, encoder)

	controller.NewAccountController(router, accountService, sessionService).AddRoutes()

	router.Use(authMiddleWare.Authorize)

	logger.Info("server starting on address " + config.Get("server.address") + ", on port " + config.Get("server.port"))
	http.ListenAndServe(config.Get("server.address")+":"+config.Get("server.port"), router)
}
