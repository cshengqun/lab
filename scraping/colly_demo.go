package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)


func CollyScraping() {
	c := colly.NewCollector()
	c.OnHTML(".info", func(e *colly.HTMLElement){
		// e.Request.Visit(e.Attr("href"))
		test := e.ChildText(".rating_num")
		fmt.Printf("test:%s\n", test)
		titla := e.ChildText(".hd a span:first-child")
		fmt.Printf("titla:%s\n", titla)
		return
		name := e.DOM.Find(".hd a").Text()
		rate := e.DOM.Find(".rating_num").Text()
		texts := strings.Split(name, "/")
		if len(texts) > 0 {
			fmt.Printf("%s %s\n", strings.TrimSpace(texts[0]), rate)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("response:", string(response.Body))
	})
	for start:=0; start<=250; start= start + 25 {
		url := fmt.Sprintf("https://movie.douban.com/top250?start=%d&filter=", start)
		c.Visit(url)
	}

}
