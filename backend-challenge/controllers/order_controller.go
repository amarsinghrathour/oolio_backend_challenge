package controllers

import (
	"fmt"
	"github.com/amarsinghrathour/oolio_backend_challenge/models"
	"github.com/amarsinghrathour/oolio_backend_challenge/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// OrderController handles all order-related requests
type OrderController struct {
	OrderService *services.OrderService
}

// NewOrderController initializes a new OrderController
func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// PlaceOrder godoc
// @Summary Place an order
// @Description Place a new order in the store
// @Tags order
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body models.OrderReq true "Order Details"
// @Success 200 {object} models.Order
// @Router /order [post]
func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	var orderReq models.OrderReq
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		log.Println("error in binding json: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Validate the promo code if it exists
	if orderReq.CouponCode != "" {
		isValid, err := ctrl.OrderService.ValidatePromoCode(orderReq.CouponCode)
		if err != nil {
			log.Println("Error validating coupon code: ", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		if !isValid {
			log.Println("Invalid coupon code")
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
	}

	// Now proceed with order placement
	order, err := ctrl.OrderService.PlaceOrder(&orderReq)
	if err != nil {
		log.Println(fmt.Sprintf("Order placement failed: %v", err))
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	// Return the successfully placed order
	c.JSON(http.StatusOK, order)
}
