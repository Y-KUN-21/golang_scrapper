package controllers

import (
	"fmt"
	models "goanimey/models"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

// ControllerDetail to get all info of given anime .
func ControllerDetail(fibGo *fiber.Ctx) {
	var danimes = make([]models.DetailAnime, 0)
	var danime models.DetailAnime = models.DetailAnime{}
	replaceTags := strings.NewReplacer("Tags", "")
	replaceDescription := strings.NewReplacer("Description", "")
	anime := fibGo.Query("anime")
	url := anime
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		danime.Name = e.ChildText("div[class=titlemobile1]")
		danime.Description = replaceDescription.Replace(e.ChildText("div[id=description-mob]"))
		danime.Tags = replaceTags.Replace(e.ChildText("div[class=tags-mobile]"))
		e.ForEach(".episodes.range.active li", func(i int, ep *colly.HTMLElement) {
			if i < 0 {
				fmt.Printf("something went wrong \n")
				return
			}
			danime.EpisodeNumber = append(danime.EpisodeNumber, ep.ChildText("a"))
			danime.EpisodeURLs = append(danime.EpisodeURLs, ep.ChildAttrs("a", "href")[0])
		})
		danimes = append(danimes, danime)
	})
	c.OnScraped(func(r *colly.Response) {
		// danimesRepo.Data = &danimes
		fibGo.Status(200).JSON(danimes)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
}
