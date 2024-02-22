package v1

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

////其他电池
func GetGetIteaqOtherDianChi(context *gin.Context) {
	err, data := models.GetIteaqOtherDianChi()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
	}
}