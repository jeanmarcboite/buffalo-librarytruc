package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeanmarcboite/librarytruc/internal/utils"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online"
)

// SearchParam -- extract search param
type SearchParam struct {
	Search string `uri:"search" binding:"required"`
}

// Search --
func Search(c *gin.Context) {
	c.Request.ParseForm()
	what := c.Request.PostForm["what"][0]

	books, err := online.SearchTitle(what)
	/*
		c.HTML(http.StatusOK, "Search.html",
			gin.H{"title": "search", "search": c.Request.PostForm["what"]})
	*/
	if err != nil {
		html := utils.JSON2HTML(err.Error())
		c.HTML(http.StatusInternalServerError, "error-500.html", gin.H{"errorHTML": html})
		return
	}

	c.HTML(http.StatusOK,
		"Search.html",
		gin.H{"title": what, "books": books, "debug": utils.SprintHTML(books)})

}
