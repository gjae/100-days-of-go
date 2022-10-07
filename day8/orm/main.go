package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func (p *Product) getPrice() uint {
	return p.Price
}

func main() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
	// Migrate schema
	db.AutoMigrate(&Product{})

	// Create record
	db.Create(&Product{Code: "0001", Price: 1500})

	// read
	var product Product
	db.First(&product) // Find product with integer primary key
	db.First(&product, "code = ?", "0001") // Find product with code 0001

	fmt.Printf("Product: %s, \nPrice: %d\n", product.Code, product.Price)

	// Update - uodate products price to 200
	db.Model(&product).Update("Price", 200)

	fmt.Printf("Product new price: %v\n", product.getPrice())
	// Update - update multiplefields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields

	// Delete delete product
	db.Delete(&product, 1)
}