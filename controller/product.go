package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/logging"
	"main/model/entity"
	"main/repository"
	"main/utils"
	"net/http"
)

// Example controller for product entity
type ProductController struct {
	ProductRepo repository.ProductRepository
	ctx         context.Context
}

// NewProductController example
func NewProductController(ctx context.Context, repo repository.ProductRepository) *ProductController {
	return &ProductController{
		ProductRepo: repo,
		ctx:         ctx,
	}
}

// AddProduct godoc
// @Summary			Add new product
// @Description 	Add new product and get entity with ID in a response
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param			product	body		entity.Product	true	"Product info"
// @Success			201	  {string}	string				"New product successfully added"
// @Failure      	400   {object}  utils.HTTPError
// @Failure      	404   {object}  utils.HTTPError
// @Failure      	500   {object}  utils.HTTPError
// @Router			/products [post]
func (c *ProductController) AddProduct(ctx *gin.Context) {
	var p entity.AddProduct
	if err := ctx.ShouldBindJSON(&p); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//TODO add struct validation
	product := entity.Product{
		Name:        p.Name,
		Description: p.Description,
		CategoryId:  p.CategoryId,
	}

	//TODO deal with contexts correctly
	c2 := context.Background()
	err := c.ProductRepo.InsertProducts(c2, []*entity.Product{
		&product,
	})
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	//TODO return struct with ID
	ctx.JSON(http.StatusOK, "Product created")
}

// ListProducts godoc
// @Summary			Get products
// @Description 	Returns all the products in system or products filtered using query
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param			query	query	string	false	"search in name, description or category"
// @Success			201		{array}		entity.Product
// @Failure      	400   	{object}  	utils.HTTPError
// @Failure      	404   	{object}  	utils.HTTPError
// @Failure      	500   	{object}  	utils.HTTPError
// @Router			/products [get]
func (c *ProductController) ListProducts(ctx *gin.Context) {
	logging.InfoFormat("List of products")
	fmt.Println("List of products")
}

// GetProduct godoc
// @Summary			Get product by ID
// @Description 	Returns product by ID
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param			id		path		int	true	"Product ID"
// @Success			200		{object}	entity.Product
// @Failure      	400   	{object}	utils.HTTPError
// @Failure      	404   	{object}	utils.HTTPError
// @Failure      	500   	{object}	utils.HTTPError
// @Router			/products/{id} [get]
func (c *ProductController) GetProduct(ctx *gin.Context) {

}

// PutProduct godoc
// @Summary			Edit product
// @Description 	Edit existing product
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param			id			path		int				true	"Product ID"
// @Param			data		body		entity.Product	true	"Product entity"
// @Success			200		{object}	entity.Product
// @Failure      	400   	{object}  	utils.HTTPError
// @Failure      	404   	{object}  	utils.HTTPError
// @Failure      	500   	{object}  	utils.HTTPError
// @Router			/products/{id} [put]
func (c *ProductController) PutProduct(ctx *gin.Context) {

}

// DeleteProduct godoc
// @Summary			Delete product
// @Description 	Delete selected product
// @Tags			Products
// @Accept			json
// @Produce			json
// @Param			id     	path	int		true	"Account ID"
// @Success			200	{string}	string	"Product deleted"
// @Failure      	400	{object}	utils.HTTPError
// @Failure      	404	{object}	utils.HTTPError
// @Failure      	500	{object}	utils.HTTPError
// @Router			/products/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {

}
