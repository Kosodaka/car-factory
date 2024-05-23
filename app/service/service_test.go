package service_test

import (
	"car-factory/app/dto"
	"car-factory/app/entity"
	"car-factory/app/service"
	mock_repo "car-factory/pkg/mocks/api/repo"
	mock_service "car-factory/pkg/mocks/api/service"
	"car-factory/pkg/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type Dependencies struct {
	repo    *mock_repo.MockCarStorage
	service *mock_service.MockCarCreator
}

func TestService_MapDto(t *testing.T) {
	car := &entity.Car{
		Brand: "Toyota",
		Color: "Grey",
		Form:  "Suv",
	}
	dtoCar := service.MapDto(car)
	require.Equal(t, dtoCar, dto.CarDto{
		Brand: car.Brand,
		Color: car.Color,
		Form:  car.Form,
	})
}

func TestService_CreateSuv(t *testing.T) {
	cases := []struct {
		name   string
		input  entity.Car
		expErr error
	}{
		{
			name: "create suv no error return",
			input: entity.Car{
				Brand: "Toyota",
				Color: "Grey",
				Form:  "Suv",
			},
			expErr: nil,
		},
	}
	ctl := gomock.NewController(t)
	d := Dependencies{
		repo: mock_repo.NewMockCarStorage(ctl),
	}
	svc := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d.repo.EXPECT().StoreCar(service.MapDto(&tt.input)).Return(nil)
			res, err := uc.CreateCar(tt.input)
			require.Equal(t, res, &tt.input)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}
func TestCreateSUV_GetCar(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		preparation func(inp string, d Dependencies) []dto.CarDto
		expErr      error
	}{
		{
			name:  "suv get car no error return",
			input: "Toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				car := dto.CarDto{
					Brand: inp,
					Color: "Grey",
					Form:  "Suv",
				}
				d.repo.EXPECT().GetCar(car.Brand).Return([]dto.CarDto{car}, nil)
				return []dto.CarDto{car}
			},
			expErr: nil,
		},
		{
			name:  "suv invalid brand",
			input: "toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				return nil
			},
			expErr: service.InvalidBrand,
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	d := Dependencies{
		repo:    mock_repo.NewMockCarStorage(ctl),
		service: mock_service.NewMockCarCreator(ctl),
	}
	svc := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cars := tt.preparation(tt.input, d)
			res, err := uc.GetCar(tt.input)
			require.Equal(t, cars, res)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}

func TestService_CreateSedan(t *testing.T) {
	cases := []struct {
		name        string
		input       entity.Car
		preparation func(inp entity.Car, d Dependencies, err error)
		expErr      error
	}{
		{
			name: "create sedan no error return",
			input: entity.Car{
				Brand: "Toyota",
				Color: "Grey",
				Form:  "Sedan",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {
				res := &entity.Car{
					Brand: "Toyota",
					Color: "Grey",
					Form:  "Sedan",
				}
				d.repo.EXPECT().StoreCar(service.MapDto(&inp)).Return(nil)
				require.Equal(t, res, &inp)
			},
			expErr: nil,
		},
		{
			name: "create sedan invalid brand",
			input: entity.Car{
				Brand: "toyota",
				Color: "Grey",
				Form:  "Sedan",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {

			},
			expErr: service.InvalidBrand,
		},
	}
	ctl := gomock.NewController(t)
	d := Dependencies{
		repo: mock_repo.NewMockCarStorage(ctl),
	}
	v := validator.NewValidator()
	svc := service.CreateSedan{}
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.preparation(tt.input, d, tt.expErr)
			_, err := uc.CreateCar(tt.input)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}
func TestCreateSedan_GetCar(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		preparation func(inp string, d Dependencies) []dto.CarDto
		expErr      error
	}{
		{
			name:  "sedan get car no error return",
			input: "Toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				car := dto.CarDto{
					Brand: inp,
					Color: "Grey",
					Form:  "Suv",
				}
				d.repo.EXPECT().GetCar(car.Brand).Return([]dto.CarDto{car}, nil)
				return []dto.CarDto{car}
			},
			expErr: nil,
		},
		{
			name:  "sedan invalid brand",
			input: "toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				return nil
			},
			expErr: service.InvalidBrand,
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	d := Dependencies{
		repo:    mock_repo.NewMockCarStorage(ctl),
		service: mock_service.NewMockCarCreator(ctl),
	}
	svc := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cars := tt.preparation(tt.input, d)
			res, err := uc.GetCar(tt.input)
			require.Equal(t, cars, res)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}

func TestService_CreateHatch(t *testing.T) {
	cases := []struct {
		name   string
		input  entity.Car
		expErr error
	}{
		{
			name: "hatchback no error return",
			input: entity.Car{
				Brand: "Toyota",
				Color: "Grey",
				Form:  "Hatchback",
			},
			expErr: nil,
		},
	}
	ctl := gomock.NewController(t)
	d := Dependencies{
		repo: mock_repo.NewMockCarStorage(ctl),
	}
	v := validator.NewValidator()
	svc := service.CreateHatchBack{}
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			d.repo.EXPECT().StoreCar(service.MapDto(&tt.input)).Return(nil)
			res, err := uc.CreateCar(tt.input)
			require.Equal(t, res, &tt.input)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}
func TestCreate_HatchGetCar(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		preparation func(inp string, d Dependencies) []dto.CarDto
		expErr      error
	}{
		{
			name:  "hatch get car no error return",
			input: "Toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				car := dto.CarDto{
					Brand: inp,
					Color: "Grey",
					Form:  "Suv",
				}
				d.repo.EXPECT().GetCar(car.Brand).Return([]dto.CarDto{car}, nil)
				return []dto.CarDto{car}
			},
			expErr: nil,
		},
		{
			name:  "hatch invalid brand",
			input: "toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				return nil
			},
			expErr: service.InvalidBrand,
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	d := Dependencies{
		repo:    mock_repo.NewMockCarStorage(ctl),
		service: mock_service.NewMockCarCreator(ctl),
	}
	svc := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cars := tt.preparation(tt.input, d)
			res, err := uc.GetCar(tt.input)
			require.Equal(t, cars, res)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}

func TestService_GetCar(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		preparation func(inp string, d Dependencies) []dto.CarDto
		expErr      error
	}{
		{
			name:  "get car no error return",
			input: "Toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				car := dto.CarDto{
					Brand: inp,
					Color: "Grey",
					Form:  "Suv",
				}
				d.repo.EXPECT().GetCar(car.Brand).Return([]dto.CarDto{car}, nil)
				return []dto.CarDto{car}
			},
			expErr: nil,
		},
		{
			name:  "invalid brand",
			input: "toyota",
			preparation: func(inp string, d Dependencies) []dto.CarDto {
				return nil
			},
			expErr: service.InvalidBrand,
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	d := Dependencies{
		repo:    mock_repo.NewMockCarStorage(ctl),
		service: mock_service.NewMockCarCreator(ctl),
	}
	svc := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, svc, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cars := tt.preparation(tt.input, d)
			res, err := uc.GetCar(tt.input)
			require.Equal(t, cars, res)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}
}

func TestService_CreateCar(t *testing.T) {
	cases := []struct {
		name        string
		input       entity.Car
		preparation func(inp entity.Car, d Dependencies, err error)
		expErr      error
	}{
		{
			name: "create car no error return",
			input: entity.Car{
				Brand: "Toyota",
				Color: "Grey",
				Form:  "Suv",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {
				d.repo.EXPECT().StoreCar(service.MapDto(&inp)).Return(err)
			},
			expErr: nil,
		},
		{
			name: "invalid Brand",
			input: entity.Car{
				Brand: "gonda",
				Color: "Grey",
				Form:  "Suv",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {},
			expErr:      service.InvalidBrand,
		},
		{
			name: "invalid Color",
			input: entity.Car{
				Brand: "Toyota",
				Color: "wHite",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {},
			expErr:      service.InvalidColor,
		},
		{
			name: "invalid Form",
			input: entity.Car{
				Brand: "Toyota",
				Color: "Grey",
				Form:  "truck",
			},
			preparation: func(inp entity.Car, d Dependencies, err error) {},
			expErr:      service.InvalidType,
		},
	}
	ctl := gomock.NewController(t)
	d := Dependencies{
		repo: mock_repo.NewMockCarStorage(ctl),
	}
	suv := service.CreateSUV{}
	v := validator.NewValidator()
	uc := service.NewCarService(d.repo, suv, v)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.preparation(tt.input, d, tt.expErr)
			_, err := uc.CreateCar(tt.input)
			if err != tt.expErr {
				t.Errorf("expected error: %v, got: %v", tt.expErr, err)
			}
		})
	}

}
