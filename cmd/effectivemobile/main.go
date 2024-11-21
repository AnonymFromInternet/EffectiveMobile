package main

import (
	"log"
	"net/http"

	configPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/config"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/dataBase/postgre"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/router"
)

func main() {
	config := configPackage.MustCreate()
	logger := logger.MustCreate(config.Mode)
	logger.Info("config was created")
	logger.Info("logger was created")

	storage := postgre.MustCreate(config.DataSourceName, logger)
	defer storage.DB.Close()
	logger.Info("connection to database was created")

	router := router.New(storage)
	logger.Info("connection to database was created")

	server := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal("package main.main: cannot start server")
	}
}
