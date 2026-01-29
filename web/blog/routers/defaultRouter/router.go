package defaultRouter

import (
	"blog/controller/defaultController"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	rg := r.Group("/")
	rg.GET("/", indexHandler)
}

func indexHandler(c *gin.Context) {
	title := defaultController.GetIndexTilte()
	cards := defaultController.GetCards()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": title,
		"cards": cards,
	})
}
