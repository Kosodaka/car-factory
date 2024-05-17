package repo

import (
	"car-factory/app/dto"
	"github.com/jmoiron/sqlx"
)

type CarStorage interface {
	StoreCar(car dto.CarDto) error
	GetCar(brand string) ([]dto.CarDto, error)
}
type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) StoreCar(car dto.CarDto) error {
	stmt := `Insert into cars (brand, color, form) values ($1, $2, $3)`
	_, err := s.db.Exec(stmt, car.Brand, car.Color, car.Form)
	if err != nil {
		return err
	}
	return nil
}
func (s Storage) GetCar(brand string) ([]dto.CarDto, error) {
	stmt := `Select * from cars where brand=$1`
	cars := []dto.CarDto{}
	if err := s.db.Select(&cars, stmt, brand); err != nil {
		return nil, err
	}
	return cars, nil
}
