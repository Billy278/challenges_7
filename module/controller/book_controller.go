package controller

import (
	"github.com/gin-gonic/gin"
)

type BookController interface {
	CreateBookCtr(ctx *gin.Context)
	UpdateBookCtr(ctx *gin.Context)
	FindByIdBookCtr(ctx *gin.Context)
	FindAllBookCtr(ctx *gin.Context)
	DeleteBookCtr(ctx *gin.Context)
}
