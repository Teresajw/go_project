package main

import (
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	//gorm.Model
	Id    int64 `gorm:"primaryKey,autoIncrement"`
	Code  string
	Price uint
}

func main() {
	// 创建连接
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	//	//DryRun: true,
	//	Logger: logger.Default.LogMode(logger.Info),
	//})

	db, err := gorm.Open(mysql.Open("root:CQGWiRshWb@tcp(192.168.112.24:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// 建表
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D89", Price: 600})

	// Read
	/*var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})*/

	// Delete - delete product
	//db.Delete(&product, 1)
}
