package controllers

import (
	"fmt"
	models "goanimey/models"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

var animes = make([]models.Anime, 0)
var anime models.Anime = models.Anime{}

// ControllerRecent for recent anime.
func ControllerRecent(fibGo *fiber.Ctx) {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)

	// On every a element which has id as recently-added attribute call callback
	c.OnHTML("div[id=recently-added]", func(e *colly.HTMLElement) {
		e.ForEach("div[id=headerDIV_4]", func(i int, e *colly.HTMLElement) {
			anime.Name = e.ChildAttr("a[id=headerA_7]", "title")
			anime.Cover = e.ChildAttr("img[id=headerIMG_6]", "src")
			link := e.ChildAttr("a[id=headerA_5]", "href")
			// passing link as parameter to ControllerRealURL function
			ControllerRealURL(fibGo, link)
			animes = append(animes, anime)

		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://4anime.to/recently-added")
}

//ControllerRealURL asda ad as
func ControllerRealURL(fibGo *fiber.Ctx, url string) {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[id=justtothetop]", func(e *colly.HTMLElement) {
		anime.Url = e.ChildAttr("a[id=titleleft]", "href")
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		fibGo.Status(200).JSON(animes)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
	})
	c.Visit(url)
}
