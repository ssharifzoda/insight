package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"insight/internal/database"
	api "insight/internal/handlers"
	"insight/internal/server"
	"insight/internal/service"
	"insight/pkg/logging"
	"insight/pkg/utils"
)

// @title Insight API SERVICE
// @version 1.0.0
// @description Api INSIGHT project

// @host localhost:5005
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := logging.GetLogger()
	if err := utils.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env value: %s", err.Error())
	}
	conn, err := database.NewPostgresGorm()
	if err != nil {
		log.Fatalf("failed to initializing db: %s", err.Error())
	}
	repository := database.NewDatabase(conn)
	services := service.NewService(repository)
	handlers := api.NewHandler(services, log)
	srv := new(server.Server)
	if err = srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
