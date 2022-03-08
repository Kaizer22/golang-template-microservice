package repository

import (
	"context"
	"main/model/entity"
)

const (
	TagInsPr     = "INSERT PRODUCTS"
	TagDelPr     = "DELETE PRODUCTS"
	TagGetAllPr  = "GET ALL PRODUCTS"
	TagGetPrById = "GET PRODUCTS BY ID"
	TagGetPrByQ  = "GET PRODUCTS BY QUERY"
)

type ProductRepository interface {
	InsertProducts(ctx context.Context, products []*entity.Product) error
	DeleteProducts(ctx context.Context, products []*entity.Product) error
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
	GetProductById(ctx context.Context, id int) (*entity.Product, error)
	GetProductsByQuery(ctx context.Context, q string) ([]*entity.Product, error)
}
