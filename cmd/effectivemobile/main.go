package main

import (
	"log"
	"net/http"

	configPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/config"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/dataBase/postgre"
	loggerPackage "github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/router"
)

func main() {
	config := configPackage.MustCreate()
	logger := loggerPackage.MustCreate(config.Mode)
	logger.Info("config was created")
	logger.Info("logger was created")

	storage := postgre.MustCreate(config.DataSourceName, logger, config.MigrationsPaths.Up)
	defer storage.DB.Close()
	logger.Info("connection to database was created")

	router := router.New(storage, config.ExternalApiUrl, logger)
	logger.Info("router was created")

	server := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
	logger.Info("server started")
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal("package main.main: cannot start server")
	}
}
