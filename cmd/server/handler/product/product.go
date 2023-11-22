package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jum8/EBE3_GoWeb.git/internal/domain"
	"github.com/jum8/EBE3_GoWeb.git/internal/product"
	
)


type Controller struct {
	service product.Service
}

func NewProductController(service product.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// List Products godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} domain.Product
// @Router /products [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	}
}

func (c *Controller) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		productFound, err := c.service.GetById(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": productFound,
		})
	}
}

func (c *Controller) HandlerSaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productRequest domain.Product
		if err := ctx.ShouldBindJSON(&productRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
			return
		}
		product, err := c.service.Save(ctx, productRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
	
}

func (c *Controller) HandlerUpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		
		var productRequest domain.Product

		if err := ctx.ShouldBindJSON(&productRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error": err,
			})
			return
		}
		product, err := c.service.Update(ctx, productRequest, idParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}

func (c *Controller) HandlerDeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		err := c.service.Delete(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product deleted",
		})
	}
}