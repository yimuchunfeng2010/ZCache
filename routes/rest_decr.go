package routes

import (
	"zCache/tool"
	"zCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestDecr(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s Decr Key:%s\n", tool.GetFileNameLine(), key)

	err := DecrBy(key,value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}
}
