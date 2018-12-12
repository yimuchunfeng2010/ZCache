package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/services"
	"ZCache/tool/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"ZCache/types"
)

func Delete(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s  Delete key: %s\n", services.GetFileNameLine(), key)

	auth, err := services.ClusterHealthCheck(types.OPERATION_TYPE_DELETE)
	if err != nil  || auth != true{
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	_, err = zdata.CoreDelete(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "status": "done"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"key": key, "status": "done"})
		return
	}
}
