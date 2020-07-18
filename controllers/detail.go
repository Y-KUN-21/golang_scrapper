package controllers

import (
	"fmt"
	models "goScraper/models"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

func ControllerDetail(fibGo *fiber.Ctx) {
	var danimes = make([]models.DetailAnime, 0)
	var danime models.DetailAnime = models.DetailAnime{}
	replacer := strings.NewReplacer("Tags", "")
	anime := fibGo.Query("anime")
	url := anime
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		danime.Name = e.ChildText("div[class=titlemobile1]")
		danime.Description = e.ChildText("div[id=description-mob] p:nth-child(2)")
		danime.Studio = e.ChildText("a[id=studiomobile1]")
		danime.Tags = replacer.Replace(e.ChildText("div[class=tags-mobile]"))

		e.ForEach(".episodes.range.active li", func(i int, ep *colly.HTMLElement) {
			if i < 0 {
				fmt.Printf("something went wrong \n")
				return
			}
			danime.EpisodesNo = append(danime.EpisodesNo, ep.ChildText("a"))
			danime.EpisodesLink = append(danime.EpisodesLink, ep.ChildAttrs("a", "href")[0])
		})
		danimes = append(danimes, danime)
	})
	c.OnScraped(func(r *colly.Response) {
		// danimesRepo.Data = &danimes
		fibGo.Status(200).JSON(danimes)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
}
