package app

import (
	"car-factory/app/controller/tui"
	"car-factory/app/repo/repo"
	"car-factory/app/service"
	"car-factory/pkg/config"
	"car-factory/pkg/logger"
	"car-factory/pkg/postgres"
	"car-factory/pkg/validator"
	"fmt"
	"github.com/c-bata/go-prompt"
	"log/slog"
	"os"
	"strings"
)

func Executor(s string) {
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
	s = strings.TrimSpace(s)
	switch s {
	case "exit", "quit":
		fmt.Println("Goodbye!")
		os.Exit(0)
		return
	case "help":
		fmt.Println("Available commands:")
		fmt.Println("1. CreateUV")
		fmt.Println("2. Create Sedan")
		fmt.Println("3. Create HatchBack")
		fmt.Println("4. Get Car")
		fmt.Println("5. Exit")
	case "Create suv":
		t.CreateSuv()
		return
	case "Create hatchback":
		t.CreateHatch()
		return
	case "Create sedan":
		t.CreateSedan()
		return
	case "Get car":
		t.GetCar()
		return
	default:
		fmt.Println("Invalid choice")
	}

}

// Плохо работате в консоле IDE лучше смотреть в терминале системы
func Completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	switch d.Text {
	case "h", "help":
		s = []prompt.Suggest{{Text: "help"}}
	case "C", "Cr", "Cre", "c", "cr", "cre":
		s = []prompt.Suggest{
			{Text: "Create", Description: "Create new car"},
			{Text: "Create hatchback"},
			{Text: "Create sedan"},
			{Text: "Create suv"},
		}
	case "G", "Get", "Get c":
		s = []prompt.Suggest{{Text: "Get car"}}
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
