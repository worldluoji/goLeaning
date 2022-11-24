package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model        // Models 是标准的 Go struct，用来代表数据库中的一个表结构。我们可以给 Models 添加 TableName 方法，来告诉 GORM 该 Models 映射到数据库中的哪张表
	Code       string `gorm:"column:code"`
	Price      uint   `gorm:"column:price"`
}

// TableName maps to mysql table name.
func (p *Product) TableName() string {
	return "product"
}

var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3309", "MariaDB service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 1. Auto migration for given models
	// GORM 的 AutoMigrate 方法，只对新增的字段或索引进行变更，理论上是没有风险的。在实际的 Go 项目开发中，也有很多人使用 AutoMigrate 方法自动同步表结构。但我更倾向于规范化、可感知的操作方式，所以我在实际开发中，都是手动变更表结构的。
	db.AutoMigrate(&Product{})

	// 2. Insert the value into database
	if err := db.Create(&Product{Code: "D42", Price: 100}).Error; err != nil {
		log.Fatalf("Create error: %v", err)
	}
	PrintProducts(db)

	// 3. Find first record that match given conditions
	product := &Product{}
	if err := db.Where("code= ?", "D42").First(&product).Error; err != nil {
		log.Fatalf("Get product error: %v", err)
	}

	// 4. Update value in database, if the value doesn't have primary key, will insert it
	product.Price = 200
	if err := db.Save(product).Error; err != nil {
		log.Fatalf("Update product error: %v", err)
	}
	PrintProducts(db)

	// 5. Delete value match given conditions
	// 这里需要注意，因为 Product 的 gorm.Model 中有 gorm.DeletedAt 字段，所以，上述删除操作不会真正把记录从数据库表中删除掉，而是通过设置数据库 product 表 deletedAt 字段为当前时间的方法来删除。
	if err := db.Where("code = ?", "D42").Delete(&Product{}).Error; err != nil {
		log.Fatalf("Delete product error: %v", err)
	}
	PrintProducts(db)
}

// List products
func PrintProducts(db *gorm.DB) {
	products := make([]*Product, 0)
	var count int64
	d := db.Where("code like ?", "%D%").Offset(0).Limit(2).Order("id desc").Find(&products).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error: %v", d.Error)
	}

	log.Printf("totalcount: %d", count)
	for _, product := range products {
		log.Printf("\tcode: %s, price: %d\n", product.Code, product.Price)
	}
}
