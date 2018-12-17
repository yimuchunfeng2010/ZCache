package routes

import (
	"ZCache/data"
	"ZCache/global"
	"ZCache/tool"
	"ZCache/tool/logrus"
	"ZCache/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(context *gin.Context) {
	key := context.Param("key")
	logrus.Infof("%s  Delete key: %s\n", tool.GetFileNameLine(), key)

	auth, err := tool.ClusterHealthCheck(types.OPERATION_TYPE_DELETE)
	if err != nil || auth != true {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	global.GlobalVar.GRWLock.Lock()
	defer global.GlobalVar.GRWLock.Unlock()
	_, err = zdata.CoreDelete(key)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "done"})
	} else {
		context.JSON(http.StatusOK, gin.H{"key": key, "status": "done"})
	}
}
