package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

//UPS
func GetIteaqUps(context *gin.Context) {
	err,data:=models.GetIteaqUps()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}