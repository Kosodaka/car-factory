package main

import (
	"car-factory/app/controller/tui"
	"car-factory/app/repo/repo"
	"car-factory/app/service"
	"car-factory/pkg/config"
	"car-factory/pkg/logger"
	"car-factory/pkg/postgres"
	"car-factory/pkg/validator"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		panic(err)
	}
	cfg := config.LoadConfig()
	log := logger.SetupLogger(cfg.GetEnv())
	log.Info("start", slog.String("env", cfg.Env))
	v := validator.NewValidator()
	psql := postgres.NewPsql(cfg.PostgresDSN)
	db, err := psql.GetDb()
	if err != nil {
		panic(err)
	}
	repository := repo.NewStorage(db)
	if err != nil {
		panic(err)
	}
	svcSuv := service.NewCarService(repository, service.CreateSUV{}, v)
	svcHatch := service.NewCarService(repository, service.CreateHatchBack{}, v)
	svcSedan := service.NewCarService(repository, service.CreateSedan{}, v)
	t := tui.NewTui(log, svcSuv, svcHatch, svcSedan)

	fmt.Println("1. Create SUV")
	fmt.Println("2. Create Sedan")
	fmt.Println("3. Create HatchBack")
	fmt.Println("4. Get Car")
	fmt.Println("5. Exit")

	for {
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			t.CreateSuv()
		case 2:
			t.CreateSedan()
		case 3:
			t.CreateHatch()
		case 4:
			t.GetCar()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}

	}

}
