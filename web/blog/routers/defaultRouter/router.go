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
	// 创建一些示例数据
	// cards := Cards{
	// 	Cards: []Card{
	// 		{
	// 			Title:   "Card 1",
	// 			Content: "This is card 1",
	// 			Date:    "2026-01-25",
	// 			Tags:    []string{"go", "gin"},
	// 		},
	// 		{
	// 			Title:   "Card 2",
	// 			Content: "This is card 2",
	// 			Date:    "2026-01-24",
	// 			Tags:    []string{"html", "template"},
	// 		},
	// 	},
	// }
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website5",
		"data": defaultController.GetCards().Cards,
	})
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	// "data": cards.Cards,
	// 	"test": "abcd",
	// })
}
