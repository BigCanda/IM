// Package service service层
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// c是context
// @Tags 首页
// @Success 200 {string} welcome //返回状态 200 返回值·
// @Router /index [get] 路径：/index 方法get
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome!",
	})
}
