package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/librarything"
	"github.com/jeanmarcboite/librarytruc/internal/utils"
	"net/http"
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
	thing, err := librarything.LookUpISBN(bookID.ISBN)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error-500.html", gin.H{"error": err})
		return
	}
	dataHTML := utils.SprintHTML(thing)

	c.HTML(http.StatusOK,
		"Details.html", gin.H{"title": "this is the title (reload)", "debug": dataHTML, "librarything": thing})
}
