package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

type usersController struct {

}

func NewUsersController() UsersController  {
	return &usersController{}
}

func (c *usersController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Create",
	})
}

func (c *usersController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Index",
	})
}

func (c *usersController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": ctx.Param("id"),
	})
}

func (c *usersController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Update",
	})
}

func (c *usersController) Destroy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Offices Destroy",
	})
}
