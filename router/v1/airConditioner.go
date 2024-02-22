package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

func GetAirConditioner1(context *gin.Context) {
	err, data := models.GetAirConditioner1()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
	}

}
func GetAirConditioner2(context *gin.Context) {
	err, data := models.GetAirConditioner2()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code":200, "msg": "success", "data": data})
	}
}
func GetAirConditioner3(context *gin.Context) {
	err, data := models.GetAirConditioner3()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code":200, "msg": "success", "data": data})
	}
}
func GetAirConditioner4(context *gin.Context) {
	err, data := models.GetAirConditioner4()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code":200, "msg": "success", "data": data})
	}
}
func GetAirConditioner5(context *gin.Context) {
	err, data := models.GetAirConditioner5()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code":200, "msg": "success", "data": data})
	}
}
func GetAirConditioner6(context *gin.Context) {
	err, data := models.GetAirConditioner6()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code":200, "msg": "success", "data": data})
	}
}