package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123asd@tcp(192.168.11.11:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 生成表结构
	_ = db.AutoMigrate(&Product{})

	// Create
	// INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2023-05-23 23:27:47.388','2023-05-23 23:27:47.388',NULL,'D42',100)
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	db.First(&product, 1)
	// SELECT * FROM `products` WHERE code = 'D42' AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	db.First(&product, "code = ?", "D42")

	// UPDATE `products` SET `price`=200,`updated_at`='2023-05-23 23:27:47.392' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Update("Price", 200)
	// updates 更新多字段
	// UPDATE `products` SET `updated_at`='2023-05-23 23:27:47.393',`code`='F42',`price`=200 WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段,比如string类型为""，int类型为0，都是不允许的, 除非字段类型为 sql.NullString...
	//UPDATE `products` SET `code`='F42',`price`=200,`updated_at`='2023-05-23 23:27:47.395' WHERE `products`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 逻辑删除 product
	// UPDATE `products` SET `deleted_at`='2023-05-23 23:27:47.397' WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL
	db.Delete(&product, 1)
}
