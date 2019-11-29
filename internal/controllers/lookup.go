package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeanmarcboite/librarytruc/internal/utils"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/librarything"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

type BookID struct {
	ISBN string `uri:"id" binding:"required"`
}

type Book struct {
	Title string
	URL   string
	Cover string
}

// LookupID --
func LookupID(c *gin.Context) {
	var bookID BookID
	if err := c.ShouldBindUri(&bookID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var work librarything.Work
	work, err := librarything.LookUpISBN(bookID.ISBN)
	if err != nil {
		html := utils.JSON2HTML(err.Error(), true)
		c.HTML(http.StatusInternalServerError, "error-500.html", gin.H{"errorHTML": html})
		return
	}
	coverURL := fmt.Sprintf(net.Koanf.String("librarything.url.cover"),
		net.Koanf.String("librarything.key"), bookID.ISBN)
	book := Book{
		Title: work.Ltml.Item.Title,
		URL:   work.Ltml.Item.URL,
		Cover: coverURL,
	}
	d := make(map[string]interface{})
	d["LibraryThing"] = work
	c.HTML(http.StatusOK,
		"Details.html",
		gin.H{"title": "this is the title (reload)",
			"book":         book,
			"debug":        utils.NewDebug(d),
			"librarything": work})
}
