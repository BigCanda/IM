package main

import (
	"Im/utils"
	"fmt"
	"gorm.io/gorm"
)

// Product 结构体在一定程度上代替了对象
type Product struct {
	// 用于数据库表的创建，方便模型定义
	// 包含了 ID, CreatedAt, UpdatedAt, DeletedAt四个属性
	gorm.Model
	ProCode  string
	ProName  string
	ProPrice uint
}

// UserBasics Golang对于类似public、private这样的访问控制并没有像Java一样严格的区分，
// 仅通过首字母的大小写来判断该字段是否可以被外部包访问。
// 在使用gorm和AutoMigration功能时，如果字段首字母为小写，
// gorm则默认该字段是私有字段，因此不会被AutoMigrate自动迁移。

func insertIntoProduct(db *gorm.DB, proCode string, proName string, proPrice uint) {
	var product Product
	err := db.Where("pro_name = ? OR pro_code = ?", proName, proCode).First(&product)
	if err.Row() == nil {
		db.Create(&Product{ProCode: proCode, ProName: proName, ProPrice: proPrice})
		return
	}
	fmt.Println("数据已存在!")
}

func updateProductPrice(db *gorm.DB, proCode string, price uint) {
	var product Product
	err := db.Model(&product).Where("pro_code = ?", proCode).Update("pro_price", price)
	if err.Row() == nil {
		fmt.Println("更新失败,商品不存在!")
		return
	}
}

func selectProductByProCode(db *gorm.DB, proCode string) Product {
	var product Product
	err := db.Where("pro_code = ?", proCode).First(&product)
	if err.Row() == nil {
		fmt.Println("该商品不存在!")
	}
	return product
}

func main() {
	//dsn := "root:111111@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("连接mysql失败")
	//}
	// 初始化执行一次 不可重复执行
	//err = DB.AutoMigrate(&UserBasic{})
	//err = DB.AutoMigrate(&models.UserBasics{})
	//if err != nil {
	//	panic("数据表创建失败！")
	//}

	//insertIntoProduct(DB, "shipin005", "宝马冰淇淋", 99999)
	//// updateProductPrice(DB, "shipin006", 20)
	// 根据整型主键查找
	//var product Product
	//product = selectProductByProCode(DB, "shipin001")
	//DB.Where("pro_name = ?", "雪碧").First(&product)
	//if product.ID >= 1 {
	//	fmt.Println("商品编码:" + product.ProCode + " 商品名:" + product.ProName + " 商品价格: " + strconv.Itoa(int(product.ProPrice)))
	//}

	// models.GetProductList()

	fmt.Println(utils.GetSalt())
}
