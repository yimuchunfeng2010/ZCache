package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sirupsen/logrus"
	"ZCache/data"
)


func Get(context *gin.Context){
	key := context.Param("key")
	logrus.Infof("Get key: %s",key)

	node , err := zdata.CoreGet(key)
	if err != nil {
		context.JSON(http.StatusNotFound,gin.H{"value":"","status":"done"})
		return
	} else {
		context.JSON(http.StatusOK,gin.H{"key":node.Data.Key,"value":node.Data.Value, "status":"done"})
		return
	}


}
