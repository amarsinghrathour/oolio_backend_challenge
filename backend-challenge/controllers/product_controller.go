// controllers/product.go
package controllers

import (
	"github.com/amarsinghrathour/oolio_backend_challenge/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProductController struct {
	Service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

// GetAllProducts ListProducts godoc
// @Summary List products
// @Description Get all products available for order
// @Tags product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Product
// @Router /product [get]
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.Service.GetAll()
	if err != nil {
		log.Println("Error getting products: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// GetProductByID GetProduct godoc
// @Summary Find product by ID
// @Description Returns a single product by ID
// @Tags product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param productId path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /product/{productId} [get]
func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productId")
	product, err := c.Service.GetByID(id)
	if err != nil {
		log.Println("Error getting product: ", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
