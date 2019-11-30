package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// Home - Home page
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "Home.html", gin.H{})
}
