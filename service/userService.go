// Package service 业务层
package service

import (
	"Im/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.JSON(-1, gin.H{
			"message": "请输入用户名!",
		})
		return
	}
	if c.Query("password") == "" {
		c.JSON(-1, gin.H{
			"message": "请输入密码!",
		})
		return
	}
	if c.Query("rePassword") == "" {
		c.JSON(-1, gin.H{
			"message": "请再次输入密码!",
		})
		return
	}

	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "两次输入密码不一致!",
		})
		return
	}
	user = models.SelectUserByUsername(username)
	if user.Username != "" {
		c.JSON(-1, gin.H{
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
// @param id formData string false "ID"
// @param password formData string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/modifyPassword [post] 路径：/user/modifyPassword 方法post
func ModifyPassword(c *gin.Context) {
	password := c.PostForm("password")
	id, _ := strconv.Atoi(c.PostForm("id"))

	if password == "" {
		c.JSON(-1, gin.H{
			"message": "请输入要修改的密码!",
		})
		return
	} else {
		models.UpdatePasswordByID(id, password)
		c.JSON(http.StatusOK, gin.H{
			"message": "密码修改成功!",
		})
		return
	}
}

// GetCode 使用gin.Context获取JSON数据
// @Summary 获取验证码
// @Tags 用户模块
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code", "message"}
// @Router /user/getCode [post] 路径：/user/getCode 方法post
func GetCode(c *gin.Context) {
	email := c.PostForm("email")
	user := models.SelectUserByEmail(email)

	if models.IsEmpty(user) {
		c.JSON(-1, gin.H{
			"message": "用户不存在!",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "邮件发送成功!",
		})
		// 发邮件
		return
	}
}

// ForgetPassword 使用gin.Context获取JSON数据
// @Summary 获取验证码后修改密码
// @Tags 用户模块
// @param email formData string false "邮箱"
// @param password formData string false "密码"
// @param rePassword formData string false "重复密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/forgetPassword [post] 路径：/user/forgetPassword 方法post
func ForgetPassword(c *gin.Context) {
	// 验证验证码是否正确
	// 施工中......
	//
	email := c.PostForm("email")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	user := models.SelectUserByEmail(email)

	if c.Query("password") == "" {
		c.JSON(-1, gin.H{
			"message": "请输入密码!",
		})
		return
	}
	if c.Query("rePassword") == "" {
		c.JSON(-1, gin.H{
			"message": "请再次输入密码!",
		})
		return
	}

	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "两次输入密码不一致!",
		})
		return
	}
	if !models.IsEmpty(user) {
		c.JSON(-1, gin.H{
			"message": "该邮箱还未注册!",
		})
		return
	}
	models.UpdatePasswordByID(int(user.ID), password)
	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功!",
	})
}
