package controllers

import (
	"fmt"
	models "goScraper/models"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

func ControllerVideo(fibGo *fiber.Ctx) {
	var videos = make([]models.Video, 0)
	var video models.Video = models.Video{}
	url := fibGo.Query("episode")
	r, _ := regexp.Compile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_+.~#?&//=]*)")
	c := colly.NewCollector(
		colly.AllowedDomains("4anime.to"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		video.Url1 = e.ChildAttr("video[id=example_video_1] source", "src")
		videoRegex := e.ChildText(".mirror-footer.cl script")
		video.Url2 = r.FindString(videoRegex)
		videos = append(videos, video)
	})

	c.OnScraped(func(r *colly.Response) {
		fibGo.Status(200).JSON(videos)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
}
