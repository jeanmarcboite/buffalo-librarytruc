package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanmarcboite/librarytruc/internal/utils"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/librarything"
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
		html := utils.JSON2HTML(err.Error(), true)
		c.HTML(http.StatusInternalServerError, "error-500.html", gin.H{"errorHTML": html})
		return
	}
	d := make(map[string]interface{})
	d["LibraryThing"] = thing
	c.HTML(http.StatusOK,
		"Details.html", 
		gin.H{"title": "this is the title (reload)", 
		"debug": utils.NewDebug(d),
		"librarything": thing})
}
