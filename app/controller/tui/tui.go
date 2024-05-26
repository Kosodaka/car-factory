package tui

import (
	"bufio"
	"car-factory/app/entity"
	"car-factory/app/service"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Tui struct {
	service []service.CarCreator
	log     *slog.Logger
}

func NewTui(log *slog.Logger, service ...service.CarCreator) *Tui {
	return &Tui{
		service: service,
		log:     log,
	}
}

func (t *Tui) CreateSuv() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите марку и цвет:")
	input, _ := reader.ReadString('\n')
	inputArr := strings.Split(input, " ")
	lastIndex := len(inputArr) - 1
	inputArr[lastIndex] = strings.TrimRight(inputArr[lastIndex], "\r\n")
	car := entity.Car{
		Brand: inputArr[0],
		Color: inputArr[lastIndex],
	}
	suv, err := t.service[0].CreateCar(car)
	if err != nil {
		t.log.Error("can't create suv", err)
		return
	}
	fmt.Println(suv)
}
func (t *Tui) CreateHatch() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите марку и цвет:")
	input, _ := reader.ReadString('\n')
	inputArr := strings.Split(input, " ")
	lastIndex := len(inputArr) - 1
	inputArr[lastIndex] = strings.TrimRight(inputArr[lastIndex], "\r\n")
	car := entity.Car{
		Brand: inputArr[0],
		Color: inputArr[lastIndex],
	}
	if inputArr[0] == "" || inputArr[lastIndex] == "" {
		t.log.Error("can't create hatch")
		return
	}

	suv, err := t.service[1].CreateCar(car)
	if err != nil {
		t.log.Error("can't create hatch", err)
		return
	}
	fmt.Println(suv)
}
func (t *Tui) CreateSedan() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите марку и цвет:")
	input, _ := reader.ReadString('\n')
	inputArr := strings.Split(input, " ")
	lastIndex := len(inputArr) - 1
	inputArr[lastIndex] = strings.TrimRight(inputArr[lastIndex], "\r\n")
	car := entity.Car{
		Brand: inputArr[0],
		Color: inputArr[lastIndex],
	}
	if inputArr[0] == "" || inputArr[lastIndex] == "" {
		t.log.Error("can't create sedan")
		return
	}
	suv, err := t.service[2].CreateCar(car)
	if err != nil {
		t.log.Error("can't create sedan", err)
		return
	}
	fmt.Println(suv)
}
func (t *Tui) GetCar() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите марку:")
	brand, err := reader.ReadString('\n')
	if err != nil {
		t.log.Error("can't read brand", err)
		return
	}
	brand = strings.TrimRight(brand, "\r\n")

	cars, err := t.service[0].GetCar(brand)
	if err != nil {
		t.log.Error("can't get cars", err)
		return
	}
	fmt.Println(cars)
}
