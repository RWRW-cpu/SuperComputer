package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

func GetIteaqPeiDianUps1s(context *gin.Context){
	err, data := models.GetIteaqPeiDianUps1s()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": data})
	}
}