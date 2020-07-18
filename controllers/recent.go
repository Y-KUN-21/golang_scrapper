package controllers

import (
	"fmt"
	models "goScraper/models"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

func ControllerRecent(fibGo *fiber.Ctx) {
	var ranimes = make([]models.Anime, 0)
	var ranime models.Anime = models.Anime{}
	r, _ := regexp.Compile("-episode.*")
	recentURLs := []string{
		"https://4anime.to/recently-added",
	}
	for index, element := range recentURLs {
		c := colly.NewCollector(
			colly.AllowedDomains("4anime.to"),
		)
		c.OnHTML("div[id=recently-added]", func(e *colly.HTMLElement) {
			e.ForEach("div[id=headerDIV_4]", func(i int, e *colly.HTMLElement) {
				ranime.Name = e.ChildText("a[id=headerA_7]")
				ranime.Url = r.ReplaceAllString(e.ChildAttr("a[id=headerA_5]", "href"), "${1}")
				ranime.ImageUrl = e.ChildAttr("img[id=headerIMG_6]", "src")
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
		fmt.Printf("%v", index)
		c.Visit(element)
	}
}
