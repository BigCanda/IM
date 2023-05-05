// Package models 数据访问层
package models

import (
	"Im/utils"
	"gorm.io/gorm"
)

type UserBasics struct {
	gorm.Model
	Username      string `valid:"required"`
	Password      string `valid:"required"`
	Salt          string
	PhoneNum      string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string `valid:"unique"`
	ClientId      string
	ClientPort    string
	DeviceInfo    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	Status        bool `default:"false"`
	IsLogout      bool
	IsAdmin       bool
}

func IsEmpty(user UserBasics) bool {
	if user.Username == "" && user.Password == "" {
		return true
	}
	return false
}

type Product struct {
	// 用于数据库表的创建，方便模型定义
	// 包含了 ID, CreatedAt, UpdatedAt, DeletedAt四个属性
	gorm.Model
	ProCode  string
	ProName  string
	ProPrice uint
}

// TableName 获取表UserBasics
func (table *UserBasics) TableName() string {
	return "user_basics"
}

// GetUserList 获取UserList
func GetUserList() []*UserBasics {
	// 初始化了一个长度为10的slice（切片）对象，每一个元素的默认值为nil，因为切片中的用户对象都将作为指针存储。
	users := make([]*UserBasics, 10)
	utils.DB.Find(&users)
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	return users
}

func GetProductList() []*Product {
	products := make([]*Product, 10)
	utils.DB.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}
	return products
}

func SelectUserByUsername(username string) UserBasics {
	user := UserBasics{}
	utils.DB.Where("username = ?", username).First(&user)
	return user
}
func SelectUserById(id int) UserBasics {
	user := UserBasics{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}
func SelectUserByPhoneNum(PhoneNum string) UserBasics {
	user := UserBasics{}
	utils.DB.Where("phone_num = ?", PhoneNum).First(&user)
	return user
}
func SelectUserByEmail(email string) UserBasics {
	user := UserBasics{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func InsertUser(user UserBasics) *gorm.DB {
	return utils.DB.Create(&user)
}

func UpdatePasswordByID(id int, password string) *gorm.DB {
	user := UserBasics{}
	utils.DB.Where("id = ?", id).First(&user)
	return utils.DB.Model(&user).Updates(UserBasics{Password: password})
}

func DeleteUserById(id int) *gorm.DB {
	user := UserBasics{}
	utils.DB.Where("id = ?", id).First(&user)
	return utils.DB.Delete(&user)
}
