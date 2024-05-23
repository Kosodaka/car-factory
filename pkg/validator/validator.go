package validator

import (
	"car-factory/app/entity"
	"car-factory/app/service"
	"regexp"
)

type Validator struct {
	re *regexp.Regexp
}

func NewValidator() *Validator {
	return &Validator{
		re: regexp.MustCompile(`^[A-Z][a-z]+$`),
	}
}

func (v Validator) ValidateBrand(brand string) error {
	if v.re.MatchString(brand) {
		return nil
	} else {
		return service.InvalidBrand
	}
}
func (v Validator) ValidateColor(color string) error {
	if v.re.MatchString(color) {
		return nil
	} else {
		return service.InvalidColor
	}
}
func (v Validator) ValidateForm(form string) error {
	if form == "Suv" || form == "Sedan" || form == "Hatchback" {
		return nil
	} else {
		return service.InvalidType
	}
}

func (v Validator) ValidateDataToCreateCar(car entity.Car) error {
	err := v.ValidateBrand(car.Brand)
	if err != nil || car.Brand == "" {
		return err
	}

	err = v.ValidateColor(car.Color)
	if err != nil || car.Color == "" {
		return err
	}
	err = v.ValidateForm(car.Form)
	if err != nil {
		return err
	}

	return nil
}

func (v Validator) ValidateDataToStoreCar(car entity.Car) error {
	err := v.ValidateBrand(car.Brand)
	if err != nil {
		return err
	}
	err = v.ValidateColor(car.Color)
	if err != nil {
		return err
	}
	return nil
}

func (v Validator) ValidateDataToGetCar(brand string) error {
	if v.re.MatchString(brand) {
		return nil
	} else {
		return service.InvalidBrand
	}
}
