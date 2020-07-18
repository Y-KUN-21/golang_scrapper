package controllers

import (
	"fmt"
	models "goScraper/models"

	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

func ControllerSearch(fibGo *fiber.Ctx) {
	var sanimes = make([]models.SearchAnime, 0)
	//var sanimesRepo models.SearchAnimeRepo = models.SearchAnimeRepo{}
	var sanime models.SearchAnime = models.SearchAnime{}
	replacer := strings.NewReplacer(" ", "+")
	anime := fibGo.Query("anime")
	searchURL := "https://4anime.to/?s=" + anime
	url := replacer.Replace(searchURL)
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("div[id=headerDIV_95]", func(i int, e *colly.HTMLElement) {
			if i > -1 {
				sanime.ImageURL = e.ChildAttr("img:nth-child(1)", "src")
				sanime.Url = e.ChildAttr("a", "href")
				sanime.Name = e.ChildText("div:nth-child(2)")
				sanime.Year = e.ChildText("span:nth-child(3)")
				sanime.Status = e.ChildText("span:nth-child(7)")
				sanime.Season = e.ChildText("span:nth-child(5)")
				sanimes = append(sanimes, sanime)
			}
			return
		})
	})
	c.OnScraped(func(r *colly.Response) {
		// sanimesRepo.Data = &sanimes
		fibGo.Status(200).JSON(sanimes)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
}
