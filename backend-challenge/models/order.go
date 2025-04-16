// models/order.go

package models

type OrderItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderReq struct {
	CouponCode string      `json:"couponCode,omitempty"`
	Items      []OrderItem `json:"items" binding:"required"`
}

// Order represents an order
// @Description Order model with list of products and coupon
type Order struct {
	CouponCode string      `json:"couponCode"`
	ID         string      `gorm:"primaryKey" json:"id"`
	Items      []OrderItem `gorm:"-" json:"items"`    // Not stored in DB directly
	Products   []Product   `gorm:"-" json:"products"` // Resolved from DB
}
