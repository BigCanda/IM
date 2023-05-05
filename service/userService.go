// Package service 业务层
package service

import (
	"Im/models"
	"Im/utils"
	"github.com/asaskevich/govalidator"
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
// @param phoneNum formData string false "电话号码"
// @param email formData string false "邮箱"
// @param username formData string false "用户名"
// @param password formData string false "密码"
// @param rePassword formData string false "重复密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/register [post] 路径：/user/register 方法post
func Register(c *gin.Context) {
	user := models.UserBasics{}
	phoneNum := c.PostForm("phoneNum")
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")

	if email == "" && phoneNum == "" {
		c.JSON(-1, gin.H{
			"message": "请输入邮箱或电话号码!",
		})
		return
	}
	if username == "" {
		c.JSON(-1, gin.H{
			"message": "请输入用户名!",
		})
		return
	}
	if password == "" {
		c.JSON(-1, gin.H{
			"message": "请输入密码!",
		})
		return
	}
	if rePassword == "" {
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
			"message": "用户名已存在!",
		})
		return
	}
	user = models.SelectUserByEmail(email)
	if user.Username != "" {
		c.JSON(-1, gin.H{
			"message": "该邮箱已被注册!",
		})
		return
	}
	user = models.SelectUserByPhoneNum(phoneNum)
	if user.Username != "" {
		c.JSON(-1, gin.H{
			"message": "该手机号已被注册!",
		})
		return
	}
	user = models.UserBasics{}
	user.Username = username
	user.Email = email
	user.PhoneNum = phoneNum
	user.Salt = utils.GetSalt()
	user.Password = utils.Md5(password + user.Salt)

	_, err := govalidator.ValidateStruct(user)
	// 发邮件或者发短信
	if err == nil {
		models.InsertUser(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请检查邮箱格式或电话号码!",
		})
	}
}

// LoginWithUsername 使用gin.Context获取JSON数据
// @Summary 使用用户名登录
// @Tags 用户模块
// @param username formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/loginWithUsername [post] 路径：/user/loginWithUsername 方法post
func LoginWithUsername(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(-1, gin.H{
			"message": "用户名不能为空!",
		})
		return
	}
	if password == "" {
		c.JSON(-1, gin.H{
			"message": "密码不能为空!",
		})
		return
	}
	user := models.SelectUserByUsername(username)
	if models.IsEmpty(user) {
		c.JSON(-1, gin.H{
			"message": "用户不存在!",
		})
		return
	}
	if utils.Md5(password+user.Salt) != user.Password {
		c.JSON(-1, gin.H{
			"message": "密码错误!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功!",
	})
	// 设置登录状态
}

// DeleteUser 使用gin.Context获取JSON数据
// @Summary 删除用户
// @Tags 用户模块
// @param id formData string false "ID"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [post] 路径：/user/deleteUser 方法post
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	if c.PostForm("id") != "" {
		models.DeleteUserById(id)
		c.JSON(http.StatusOK, gin.H{
			"message": "用户删除成功!",
		})
	} else {
		c.JSON(-1, gin.H{
			"message": "用户删除失败,id不能为空!",
		})
	}
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
