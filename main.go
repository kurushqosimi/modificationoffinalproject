package main

import (
	"fmt"
	"log"
	"net/http"
	"one-day-job/config"
	one_day_job "one-day-job/internal/adapter/one-day-job"
	"one-day-job/internal/db/pg"
	"one-day-job/internal/domain/Kurush"
	Kurush2 "one-day-job/internal/handler/Kurush"
	"one-day-job/internal/server"
	"one-day-job/internal/storage"
	"one-day-job/pkg/logger"
)

func main() {
	cfg, err := config.InitConfigs()
	if err != nil {
		log.Fatalf("failed to init configs %v", err)
	}

	db := pg.InitDB(&cfg.Database)
	appLogger := logger.InitLogger(&cfg.Logger)
	repo := storage.NewStorage(db, appLogger)
	adapters := one_day_job.NewAdapter(&cfg.Adapter, appLogger, &http.Client{})
	services := Kurush.NewService(adapters, repo)
	handlers := Kurush2.NewHandler(services, cfg, appLogger)

	srv := server.NewServer(cfg, handlers)
	fmt.Println("application started...")
	srv.Run()

}
