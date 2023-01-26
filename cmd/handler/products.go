package handler

import (
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/gin-gonic/gin"
)

type Products struct {
	s products.Service
}

func NewHandlerProducts(s products.Service) *Products {
	return &Products{s}
}

func (p *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, products)
	}
}

func (p *Products) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := domain.Product{}
		err := ctx.ShouldBindJSON(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = p.s.Create(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": products})
	}
}

func (p *Products) LoadProductData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := p.s.LoadData()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": "Data cargada exitosamente"})
	}
}
func (p *Products) Top5Products() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Top5Products, err := p.s.Top5Products()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, Top5Products)
	}
}
