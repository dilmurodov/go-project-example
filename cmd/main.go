package main

import (
	"context"
	"project/api"
	"project/api/handlers"
	"project/config"
	"project/pkg/logger"
	"project/storage/postgres"
)

func main() {
	
	cfg := config.Load()

	log := logger.NewLogger("project-user", "debug")
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil{
		log.Panic("!!!Error while connecting to Postgres", logger.Error(err))
	}
	defer pgStore.CloseDB()
	log.Info("!!!Successfully connected to DB", logger.Any("port", cfg.PostgresPort))

	h := handlers.NewHandler(cfg, log, pgStore)
	r := api.SetUpRouter(h, cfg)

	err = r.Run(cfg.ServerPort)
	if err != nil{
		log.Error("!!!HTTP: Error server failed", logger.Error(err))
		return
	}
	log.Info("!!!HTTP: Server Being started...", logger.Any("port", cfg.ServerPort))
}