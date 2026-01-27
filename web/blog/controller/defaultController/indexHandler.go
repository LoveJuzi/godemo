package defaultController

type Card struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Date    string   `json:"date"`
	Tags    []string `json:"tags"`
}

type Cards struct {
	Cards []Card
}

func GetCards() Cards {
	// 创建一些示例数据
	cards := Cards{
		Cards: []Card{
			{
				Title:   "Card 1",
				Content: "This is card 1",
				Date:    "2026-01-25",
				Tags:    []string{"go", "gin"},
			},
			{
				Title:   "Card 2",
				Content: "This is card 2",
				Date:    "2026-01-24",
				Tags:    []string{"html", "template"},
			},
		},
	}
	return cards
}

//func IndexHandler(c *gin.Context) {
//	// 创建一些示例数据
//	// cards := Cards{
//	// 	Cards: []Card{
//	// 		{
//	// 			Title:   "Card 1",
//	// 			Content: "This is card 1",
//	// 			Date:    "2026-01-25",
//	// 			Tags:    []string{"go", "gin"},
//	// 		},
//	// 		{
//	// 			Title:   "Card 2",
//	// 			Content: "This is card 2",
//	// 			Date:    "2026-01-24",
//	// 			Tags:    []string{"html", "template"},
//	// 		},
//	// 	},
//	// }
//	c.HTML(http.StatusOK, "index.html", gin.H{
//		"title": "Main website3",
//	})
//	// c.HTML(http.StatusOK, "index.html", gin.H{
//	// 	// "data": cards.Cards,
//	// 	"test": "abcd",
//	// })
//}
