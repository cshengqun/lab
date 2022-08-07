package scraping

import (
	"encoding/json"
	"fmt"
	"github.com/cshengqun/lab/scraping/value_object"
	"github.com/gocolly/colly"
	"strings"
)


func CollyGetHtml() {
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

func CollyPostBE() {
	var question []string
	res := &value_object.BdQuestion{}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL)
	})
	c.OnResponse(func(response *colly.Response) {
		// fmt.Printf("res: %s\n", string(response.Body))
		err := json.Unmarshal(response.Body, res)
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}
		//fmt.Printf("%v\n", res)
		//fmt.Printf("plus:%+v\n", res)
		for _, item := range res.Data.Items {
			question = append(question, item.Basic.Name[0].Text)
		}

	})
	c.OnScraped(func(response *colly.Response) {
		/*
			for i, q := range question {
				fmt.Printf("%d.%s\n", i, q)
			}
		*/
		//fmt.Printf("on scraped\n")
	})
}
