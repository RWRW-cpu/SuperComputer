package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

//CPU
func GetGatewayCpuUsage(context *gin.Context) {
	err,data:=models.GetGatewayCpuUsage()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}

//硬盘
func GetGatewayDiskUsage(context *gin.Context) {
	err,data:=models.GetGatewayDiskUsage()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}

//内存
func GetGatewayMemoryUsage(context *gin.Context) {
	err,data:=models.GetGatewayMemoryUsage()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}

//登录
func GetGatewaySession(context *gin.Context) {
	err,data:=models.GetGatewaySession()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}