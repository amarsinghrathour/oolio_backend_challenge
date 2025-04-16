package services

import (
	"fmt"
	"github.com/amarsinghrathour/oolio_backend_challenge/models"
	"github.com/amarsinghrathour/oolio_backend_challenge/utility"
	"github.com/google/uuid"
)

// OrderService handles the business logic for orders.
type OrderService struct {
	productService ProductService // Inject ProductService
}

// NewOrderService creates a new instance of OrderService with injected ProductService.
func NewOrderService(productService ProductService) *OrderService {
	return &OrderService{productService: productService}
}

// PlaceOrder handles placing an order
func (s *OrderService) PlaceOrder(orderReq *models.OrderReq) (*models.Order, error) {
	orderID := uuid.New().String() // In real application, generate this using a UUID or other mechanism

	// Create Order items from the request
	orderItems := orderReq.Items

	// Fetch the products from the database based on order items
	var products []models.Product
	for _, item := range orderItems {
		// Use ProductService's GetByID method to fetch the product by ID
		product, err := s.productService.GetByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product with id %s not found", item.ProductID)
		}
		products = append(products, *product)
	}

	// Create and return the order response
	order := &models.Order{
		CouponCode: orderReq.CouponCode,
		ID:         orderID,
		Items:      orderItems,
		Products:   products,
	}

	return order, nil

}

// ValidatePromoCode validates if the given promo code is valid
func (s *OrderService) ValidatePromoCode(code string) (bool, error) {
	return utility.ValidatePromoCode(code)
}
