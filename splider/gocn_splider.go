package splider

import (
	"fmt"
	"news/base"
	"news/config"
	"news/db"
	"news/message"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var c *colly.Collector

func Run() {
	url := ""
	// On every a element which has href attribute call callback
	c.OnHTML("div[class=\"aw-mod aw-question-detail aw-item\"]", func(e *colly.HTMLElement) {
		dailyTitle := e.ChildText("h1")
		fmt.Printf("dailyTitle:%v\n", dailyTitle)
		var textUrls []db.MsgTextUrl
		var author string
		e.ForEach("div[class=\"content markitup-box\"]", func(i int, e *colly.HTMLElement) {
			e.ForEach("li", func(i int, e *colly.HTMLElement) {
				url := e.ChildText("a[href]")
				urlIndex := strings.Index(e.Text, url)
				if urlIndex <= len(e.Text) && urlIndex != -1 {
					//println(urlIndex, len(e.Text), e.Text)
					text := e.Text[0:urlIndex]
					textUrls = append(textUrls, db.MsgTextUrl{MsgText: text, MsgUrl: url})
				}
			})

			e.ForEach("*", func(i int, element *colly.HTMLElement) {
				authorIndex1 := strings.Index(element.Text, "编辑:")
				authorIndex2 := strings.Index(element.Text, "编辑：")

				index := -1
				authorLen := -1
				if authorIndex1 >= 0 {
					index = authorIndex1
					authorLen = len("编辑:")
				} else if authorIndex2 >= 0 {
					index = authorIndex2
					authorLen = len("编辑：")
				}

				if index >= 0 {
					author = element.Text[index+authorLen:]
				}
			})
		})

		message.Push(db.Message{
			DailyTitle: dailyTitle,
			TextUrls:   textUrls,
			Author:     author,
			PostUrl:    e.Request.URL.String(),
		})
	})

	c.OnHTML("div[class=\"aw-question-content\"]", func(element *colly.HTMLElement) {
		//fmt.Printf("e:%v\n",element)
		element.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			//"a[href]", func(e *colly.HTMLElement) {
			link := element.Attr("href")
			//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			if /*e.Text == ">" ||*/ strings.Contains(element.Text, "GoCN每日新闻") {
				// Visit link found on page
				// Only those links are visited which are in AllowedDomains
				//fmt.Printf("e:%v %v\n", element, element.Request.AbsoluteURL(link))
				if strings.Contains(element.Text, time.Now().Format("2006-01-02")) {
					fmt.Printf("link:%v\n", link)
					_ = c.Visit(element.Request.AbsoluteURL(link))
				}
				return
				url = element.Request.AbsoluteURL(link)
			}
		})
	})

	if config.Cfg().Splider.All {
		c.OnHTML("div[class=\"page-control\"]", func(element *colly.HTMLElement) {
			isHasNext := false
			element.ForEach("a[href]", func(i int, e *colly.HTMLElement) {
				link := e.Attr("href")
				if link == "javascript:;" {
					isHasNext = true
				}

				if isHasNext == true && link != "javascript:;" {
					_ = c.Visit(e.Request.AbsoluteURL(link))
					isHasNext = false
				}
			})
		})
	}

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	_ = c.Visit("https://gocn.vip/explore/category-14")
}

func init() {
	// Instantiate default collector
	c = colly.NewCollector(
		colly.AllowedDomains("gocn.vip"),
		colly.MaxDepth(1),
		//colly.Debugger(&debug.base.LogDebugger{}),
	)

	parallelism := config.Cfg().Splider.Parallelism
	delay := time.Duration(config.Cfg().Splider.Delay) * time.Second

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "gocn.vip",
		Parallelism: parallelism,
		Delay:       delay,
	})
	base.Log("c.limit error:%v", err)
}
