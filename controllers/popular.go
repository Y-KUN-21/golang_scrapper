package controllers

import (
	"fmt"
	models "goanimey/models"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

// ControllerPopular for popular anime.
func ControllerPopular(fibGo *fiber.Ctx) {
	var animes = make([]models.Anime, 0)
	var anime models.Anime = models.Anime{}
	urls := []string{
		"https://4anime.to/popular-this-week",
		"https://4anime.to/popular-this-week/page/2",
		"https://4anime.to/popular-this-week/page/3",
		"https://4anime.to/popular-this-week/page/4",
	}

	for index, element := range urls {
		c := colly.NewCollector(
			colly.AllowedDomains("4anime.to"),
		)

		c.OnHTML("div[id=popular-this-week]", func(e *colly.HTMLElement) {
			e.ForEach("div[id=headerDIV_4]", func(i int, e *colly.HTMLElement) {
				anime.Name = e.ChildAttr("a[id=headerA_7]", "title")
				anime.Url = e.ChildAttr("a[id=headerA_5]", "href")
				anime.Cover = e.ChildAttr("img[id=headerIMG_6]", "src")
				animes = append(animes, anime)
			})
		})
		c.OnScraped(func(r *colly.Response) {
			fibGo.Status(200).JSON(animes)
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
