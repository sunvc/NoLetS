package controller

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunvc/NoLets/common"
	"github.com/sunvc/NoLets/database"
)

func Info(c *gin.Context) {
	admin, ok := c.Get("admin")
	system := common.LocalConfig.System

	results := gin.H{
		"version": system.Version,
		"build":   system.BuildDate,
		"commit":  system.CommitID,
	}

	if ok && admin.(bool) {
		devices, _ := database.DB.CountAll()
		results["devices"] = devices
		results["arch"] = runtime.GOOS + "/" + runtime.GOARCH
		results["cpu"] = runtime.NumCPU()
	}
	c.JSON(http.StatusOK, results)
}

// Ping 处理心跳检测请求
// 返回服务器当前状态
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, common.BaseResp{
		Code:      http.StatusOK,
		Message:   "pong",
		Timestamp: time.Now().Unix(),
	})
}

func Health(c *gin.Context) { c.String(http.StatusOK, "OK") }
