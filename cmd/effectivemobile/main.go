package main

import (
	"github.com/AnonymFromInternet/EffectiveMobile/internal/config"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/dataBase/postgre"
	"github.com/AnonymFromInternet/EffectiveMobile/internal/logger"
)

func main() {
	// get config
	config := config.MustCreate()
	// get logger
	logger := logger.MustCreate(config.Mode)
	logger.Info("config was created")
	logger.Info("logger was created")
	// connect to database
	storage := postgre.MustCreate(config.DataSourceName, logger)
	logger.Info("connection to db was created")
	_ = storage
	// create routing
	// start server
}
