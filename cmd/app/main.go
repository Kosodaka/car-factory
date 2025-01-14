package main

import (
	v1 "car-factory/app/controller/http/v1"
	"car-factory/app/repo/repo"
	"car-factory/app/service"
	"car-factory/pkg/config"
	"car-factory/pkg/logger"
	"car-factory/pkg/sqlite"
	"log/slog"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	log := logger.SetupLogger(cfg.GetEnv())
	log.Info("start", slog.String("env", cfg.Env))
	sqlite := sqlite.NewSqlite(cfg.SqlitePath)
	db, err := sqlite.GetDB()
	if err != nil {
		panic(err)
	}
	repository := repo.NewStorage(db)
	err = repository.CreateTable()
	if err != nil {
		panic(err)
	}
	svcSuv := service.NewCarService(repository, service.CreateSUV{})
	svcHatch := service.NewCarService(repository, service.CreateHatchBack{})
	svcSedan := service.NewCarService(repository, service.CreateSedan{})
	server := v1.NewRouter(log, svcSuv, svcHatch, svcSedan)
	err = server.Run(cfg.HttpHost + ":" + cfg.HttpPort)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}
