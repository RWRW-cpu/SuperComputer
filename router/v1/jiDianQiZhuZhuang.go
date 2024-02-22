package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

//继电器柱状图
func GetIteaqJiDianQiZhuZhuang(context *gin.Context) {
	err, data := models.GetIteaqJiDianQiZhuZhuang()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
	}
}