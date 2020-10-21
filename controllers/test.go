package controllers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

// ControllerTest for test gogoanime.
func ControllerTest(fibGo *fiber.Ctx) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(),
	)
	var link string

	// On every a element which has href attribute call callback
	c.OnHTML(".anime_muti_link", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			link = el.ChildAttr("a", "data-video")
			fmt.Printf("Link found: %s\n\n", link)

		})
		c.Visit(link)
		c.OnHTML("#player", func(mp *colly.HTMLElement) {
			vl := e.ChildAttr("video", "src")
			fmt.Printf("video found: %s\n", vl)
		})

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://gogoanime.so/black-clover-tv-episode-148")
}
