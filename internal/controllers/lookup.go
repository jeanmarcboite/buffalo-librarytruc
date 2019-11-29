package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeanmarcboite/librarytruc/internal/utils"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online"
)

type BookID struct {
	ISBN string `uri:"id" binding:"required"`
}

// LookupID --
func LookupID(c *gin.Context) {
	var bookID BookID
	if err := c.ShouldBindUri(&bookID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	book, err := online.LookUpISBN(bookID.ISBN)

	if err != nil {
		html := utils.JSON2HTML(err.Error(), true)
		c.HTML(http.StatusInternalServerError, "error-500.html", gin.H{"errorHTML": html})
		return
	}

	c.HTML(http.StatusOK,
		"Details.html",
		gin.H{"title": book.Title, "book": book, "debug": utils.SprintHTML(book, false)})
}
