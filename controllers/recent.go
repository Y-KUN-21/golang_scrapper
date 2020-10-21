package controllers

import (
	"fmt"
	models "goanimey/models"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

// ControllerRecent for recent anime.
func ControllerRecent(fibGo *fiber.Ctx) {
	r, _ := regexp.Compile("/image/(.*)")
	var animes = make([]models.Anime, 0)
	var anime models.Anime = models.Anime{}
	urls := []string{
		"https://4anime.to/recently-added",
		"https://4anime.to/recently-added/page/2",
		"https://4anime.to/recently-added/page/3",
		"https://4anime.to/recently-added/page/4",
	}

	for index, element := range urls {
		c := colly.NewCollector(
			colly.AllowedDomains("4anime.to"),
		)

		c.OnHTML("div[id=recently-added]", func(e *colly.HTMLElement) {
			e.ForEach("div[id=headerDIV_4]", func(i int, e *colly.HTMLElement) {
				anime.Name = e.ChildAttr("a[id=headerA_7]", "title")
				anime.Url = e.ChildAttr("a[id=headerA_5]", "href")
				anime.Cover = e.ChildAttr("img[id=headerIMG_6]", "src")
				fmt.Println(r.FindString(anime.Cover))
				animes = append(animes, anime)
			})
		})
		c.OnScraped(func(r *colly.Response) {
			fibGo.Status(200).JSON(animes)
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})
		fmt.Printf("%v", index)
		c.Visit(element)
	}
}
