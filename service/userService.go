// Package service 业务层
package service

import (
	"Im/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList 使用gin.Context获取JSON数据
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get] 路径：/user/getUserList 方法get
func GetUserList(c *gin.Context) {
	users := make([]*models.UserBasics, 10)
	users = models.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}

// GetProductList 使用gin.Context获取JSON数据
// c是context
// @Tags 商品列表
// @Success 200 {string} json{"code", "message"}
// @Router /product [get] 路径：/product 方法get
func GetProductList(c *gin.Context) {
	products := make([]*models.Product, 10)
	products = models.GetProductList()

	c.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}

// Register 使用gin.Context获取JSON数据
// @Summary 新增用户
// @Tags 用户模块
// @param username query string false "用户名"
// @param password query string false "密码"
// @param rePassword query string false "重复密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/register [get] 路径：/user/register 方法get
func Register(c *gin.Context) {
	user := models.UserBasics{}
	username := c.Query("username")
	password := c.Query("password")
	rePassword := c.Query("rePassword")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入用户名!",
		})
		return
	}
	if c.Query("password") == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入密码!",
		})
		return
	}
	if c.Query("rePassword") == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "请再次输入密码!",
		})
		return
	}

	if password != rePassword {
		c.JSON(http.StatusOK, gin.H{
			"message": "两次输入密码不一致!",
		})
		return
	}
	user = models.SelectUserByUsername(username)
	if user.Username != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户已存在!",
		})
		return
	}
	user = models.UserBasics{}
	user.Username = c.Query("username")
	user.Password = password
	models.InsertUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "用户添加成功!",
	})
}

// ModifyPassword 使用gin.Context获取JSON数据
// @Summary 修改用户密码
// @Tags 用户模块
// @param password query string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/modifyPassword [get] 路径：/user/modifyPassword 方法get
func ModifyPassword(c *gin.Context) {
	user := models.UserBasics{}
	password := c.Query("password")

	if password == "" {
		c.JSON(-1, gin.H{
			"message": "请输入要修改的密码!",
		})
	} else {
		models.UpdatePassword(user, password)
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入要修改的密码!",
		})
	}
}
