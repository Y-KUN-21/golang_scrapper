package controllers

import (
	"fmt"
	models "goanimey/models"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

// ControllerRandom for random anime.
func ControllerRandom(fibGo *fiber.Ctx) {
	var ranimes = make([]models.Anime, 0)
	var ranime models.Anime = models.Anime{}
	random := fibGo.Query("page")
	// r, _ := regexp.Compile("-episode.*")
	url := fmt.Sprintf("https://4anime.to/page/%v", random)
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)
	c.OnHTML("div[id=randomani2]", func(e *colly.HTMLElement) {
		e.ForEach("div[id=headerDIV_4]", func(i int, e *colly.HTMLElement) {
			ranime.Name = e.ChildAttr("a[id=headerA_7]", "title")
			ranime.Url = e.ChildAttr("a[id=headerA_7]", "href")
			ranime.Cover = "https://4anime.to/" + e.ChildAttr("div[id=aroundimage] img", "src")
			ranimes = append(ranimes, ranime)

		},
		)
	})
	c.OnScraped(func(r *colly.Response) {
		fibGo.Status(200).JSON(ranimes)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

}
