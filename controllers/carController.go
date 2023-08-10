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

	if arr := ctx.ShouldBindJSON(&newCar); err != nil{
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
}