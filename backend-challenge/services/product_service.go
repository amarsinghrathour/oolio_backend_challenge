package services

import (
	"github.com/amarsinghrathour/oolio_backend_challenge/models"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{db: db}
}

func (s *productService) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := s.db.Find(&products).Error
	return products, err
}

func (s *productService) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := s.db.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
