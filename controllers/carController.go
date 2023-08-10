package controller

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct{
	CarId string	`json:"car_id"`
	Brand string	`json:"car_brand"`
	Model string	`json:"car_model"`
	Price int		`json:"car_price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context){
	var newCar	Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car" : newCar
	})
}

func UpdateCar(ctx *gin.Context){
	carID := ctx.Param("carID")
	condition := false
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas{
		if carID == car.CarID{
			condition = true
			CarDatas[i] = updateCar
			CarDatas[i].carID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status" : "Page Not Found",
			"error_message" : fmt.Sprintf("car with id %v not found", carID)
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}