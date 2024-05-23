package service

import (
	"car-factory/app/dto"
	"car-factory/app/entity"
	"car-factory/app/repo/repo"
	"errors"
)

var (
	InvalidType  = errors.New("wrong type, please choose correct type")
	InvalidBrand = errors.New("invalid brand")
	InvalidColor = errors.New("invalid color")
)

type Validator interface {
	ValidateColor(color string) error
	ValidateBrand(brand string) error
	ValidateForm(form string) error
	ValidateDataToCreateCar(car entity.Car) error
	ValidateDataToStoreCar(car entity.Car) error
	ValidateDataToGetCar(brand string) error
}

type CarCreator interface {
	CreateCar(req entity.Car) (*entity.Car, error)
	GetCar(brand string) ([]dto.CarDto, error)
}

type CarService struct {
	carCreator CarCreator
	carStorage repo.CarStorage
	validator  Validator
}

func NewCarService(carStorage repo.CarStorage, carCreator CarCreator, v Validator) *CarService {
	return &CarService{
		carCreator: carCreator,
		carStorage: carStorage,
		validator:  v,
	}
}

func MapDto(car *entity.Car) dto.CarDto {
	return dto.CarDto{
		Brand: car.Brand,
		Color: car.Color,
		Form:  car.Form,
	}
}

type CreateSUV struct {
	CarService
}

func (suv CreateSUV) CreateCar(req entity.Car) (*entity.Car, error) {
	car := &entity.Car{
		Brand: req.Brand,
		Color: req.Color,
		Form:  "Suv",
	}
	return car, nil
}
func (suv CreateSUV) GetCar(brand string) ([]dto.CarDto, error) {
	car, err := suv.carStorage.GetCar(brand)
	if err != nil {
		return nil, err
	}
	return car, nil
}

type CreateSedan struct {
	CarService
}

func (s CreateSedan) CreateCar(req entity.Car) (*entity.Car, error) {
	car := &entity.Car{
		Brand: req.Brand,
		Color: req.Color,
		Form:  "Sedan",
	}

	return car, nil
}

func (s CreateSedan) GetCar(brand string) ([]dto.CarDto, error) {
	car, err := s.carStorage.GetCar(brand)
	if err != nil {
		return nil, err
	}
	return car, nil
}

type CreateHatchBack struct {
	CarService
}

func (h CreateHatchBack) CreateCar(req entity.Car) (*entity.Car, error) {
	car := &entity.Car{
		Brand: req.Brand,
		Color: req.Color,
		Form:  "Hatchback",
	}

	return car, nil
}

func (h CreateHatchBack) GetCar(brand string) ([]dto.CarDto, error) {
	car, err := h.carStorage.GetCar(brand)
	if err != nil {
		return nil, err
	}
	return car, nil
}
func (cs *CarService) CreateCar(req entity.Car) (*entity.Car, error) {
	if err := cs.validator.ValidateDataToCreateCar(req); err != nil {
		return nil, err
	}
	car, err := cs.carCreator.CreateCar(req)
	if err != nil {
		return nil, err
	}
	if err = cs.validator.ValidateDataToStoreCar(*car); err != nil {
		return nil, err
	}
	err = cs.carStorage.StoreCar(MapDto(car))
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (cs *CarService) GetCar(brand string) ([]dto.CarDto, error) {
	err := cs.validator.ValidateDataToGetCar(brand)
	if err != nil {
		return nil, err
	}
	cars, err := cs.carStorage.GetCar(brand)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
