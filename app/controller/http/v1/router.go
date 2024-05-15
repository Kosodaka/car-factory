package v1

import (
	"car-factory/app/entity"
	"car-factory/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Routes struct {
	log     *slog.Logger
	service []service.CarCreator
}

func NewRouter(log *slog.Logger, service ...service.CarCreator) *gin.Engine {
	r := &Routes{
		log:     log,
		service: service,
	}
	handler := gin.Default()
	handler.POST("/create-suv", r.CreateSuv)
	handler.POST("/create-sedan", r.CreateSedan)
	handler.POST("/create-hatch", r.CreateHatchBack)

	handler.GET("/get-car/:brand", r.GetCar)
	return handler
}

func (r *Routes) CreateSuv(c *gin.Context) {
	input := entity.Car{}

	if err := c.ShouldBindJSON(&input); err != nil {
		r.log.Error(err.Error())
		fmt.Println("invalid json bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := r.service[0].CreateCar(input)
	if err != nil {
		fmt.Println("error to create car with calling to service")
		r.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}
func (r *Routes) CreateSedan(c *gin.Context) {
	input := entity.Car{}

	if err := c.ShouldBindJSON(&input); err != nil {
		r.log.Error(err.Error())
		fmt.Println("invalid json bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := r.service[2].CreateCar(input)
	if err != nil {
		fmt.Println("error to create car with calling to service")
		r.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}
func (r *Routes) CreateHatchBack(c *gin.Context) {
	input := entity.Car{}

	if err := c.ShouldBindJSON(&input); err != nil {
		r.log.Error(err.Error())
		fmt.Println("invalid json bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := r.service[1].CreateCar(input)
	if err != nil {
		fmt.Println("error to create car with calling to service")
		r.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}
func (r *Routes) GetCar(c *gin.Context) {
	brand := c.Param("brand")

	car, err := r.service[0].GetCar(brand)
	if err != nil {
		r.log.Error(err.Error())
		fmt.Println("error to call service.getcar")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, car)
}
