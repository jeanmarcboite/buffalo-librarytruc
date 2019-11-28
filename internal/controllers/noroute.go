package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error-404.html", gin.H{})
}
