package internal

import "github.com/gocolly/colly/v2"

func Parse(URL string) ([]Article, error) {
	var articles []Article
	c := colly.NewCollector(colly.Async(true))
	c.OnHTML("a.list-item__title", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")

		articles = append(articles, Article{Title: title, Link: link, Source: "ria"})
	})
	c.Visit(URL)
	c.Wait()
	return articles, nil
}
