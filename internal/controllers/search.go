package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchParam -- extract search param
type SearchParam struct {
	Search string `uri:"search" binding:"required"`
}

// Search --
func Search(c *gin.Context) {
	c.Request.ParseForm()
	for key, value := range c.Request.PostForm {
		fmt.Println(key, value)
	}
	c.HTML(http.StatusOK, "Search.html",
		gin.H{"title": "search", "search": c.Request.PostForm["what"]})
}
