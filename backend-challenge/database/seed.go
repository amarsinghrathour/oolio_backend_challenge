package database

import (
	"github.com/amarsinghrathour/oolio_backend_challenge/models"
	"gorm.io/gorm"
	"log"
)

func SeedProducts(db *gorm.DB) {
	// Check if already seeded
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("Products already seeded.")
		return
	}

	products := []models.Product{
		{ID: "10", Name: "Chicken Waffle", Price: 120.0, Category: "Waffle"},
		{ID: "11", Name: "Paneer Roll", Price: 80.0, Category: "Roll"},
		{ID: "12", Name: "Veg Burger", Price: 90.0, Category: "Burger"},
		{ID: "13", Name: "Fries", Price: 60.0, Category: "Sides"},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Println("Failed to seed products:", err)
		return
	}

	log.Println("Sample products seeded successfully.")
}
