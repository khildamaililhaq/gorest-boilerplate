package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OfficesController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

type officesController struct {

}

func NewOfficesController() OfficesController  {
	return &officesController{}
}

func (c *officesController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Create",
	})
}

func (c *officesController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Index",
	})
}

func (c *officesController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": ctx.Param("id"),
	})
}

func (c *officesController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Update",
	})
}

func (c *officesController) Destroy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Destroy",
	})
}
