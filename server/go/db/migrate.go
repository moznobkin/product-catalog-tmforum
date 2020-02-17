package db

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	sw "product-catalog-tmforum/server/swagger"
)

func migrateCatalog() {
	db, err := gorm.Open("", "product catalog")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()
  
	// Migrate the schema
	db.AutoMigrate(&sw.Catalog{})
  
	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})
  
	// // Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
  
	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)
  
	// // Delete - delete product
	// db.Delete(&product)
  }
