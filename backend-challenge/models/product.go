// models/product.go
package models

// Product represents a product in the store
// @Description Product model with ID, Name, Price, Category
type Product struct {
	ID       string  `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
