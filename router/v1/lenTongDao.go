package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

//冷通道烟感
func GetLenTongDaoYanGan(context *gin.Context) {
	err,data:=models.LenTongDaoYanGan()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}

}
//冷通道温度
func GetLenTongDaoWenDu(context *gin.Context){
	err,data:=models.GetIteaqLengTongDaoWenDu()
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{"code":404,"msg":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"msg":"success","data":data})
	}
}

