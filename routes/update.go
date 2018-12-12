package routes

import (
	"ZCache/data"
	"ZCache/types"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(context *gin.Context) {
	auth, err := services.ClusterHealthCheck(types.OPERATION_TYPE_POST)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("%s Update Key:%s, Value:%s\n", services.GetFileNameLine(), key, value)
	node, err := zdata.CoreUpdate(key, value)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"key": key, "value": value, "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": node.Key, "value": node.Value, "status": "done"})
		return
	}
}
