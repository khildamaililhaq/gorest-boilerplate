package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthsController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type authsController struct {

}

func NewAuthsController() AuthsController  {
	return &authsController{}
}

func (c *authsController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Login",
	})
}

func (c *authsController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Register",
	})
}

func (c *authsController) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"Hello Me",
	})
}
