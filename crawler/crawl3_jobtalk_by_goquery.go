package main

/**
 *	勉強用にクロールするもの作ってみた
 * 	今後はこれをDBに入れるところまで作ってみようかな
 */

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Url  string
	star string
}

const (
	PAGE_LIMIT = 1000
)

var results = []Result{}

func GetPage(url string) {

	doc, e := goquery.NewDocument(url)
	if e == nil {
		var result Result
		doc.Find("div.listBox").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.ChildrenFiltered("div.ttl").Find("a").Attr("href")
			star := s.ChildrenFiltered("div.jobContent").ChildrenFiltered("div.txt").Find("em").Text()

			//star := 4.0
			//fmt.Println(s.ChildrenFiltered("div.jobContent div.txt").Text())
			result.Url = url
			result.star = star
			results = append(results, result)
		})

	}

}

func main() {
	url := "http://jobtalk.jp/company/index_1.html"
	GetPage(url)

	for _, result := range results {
		fmt.Println(result.Url)
		fmt.Println(result.star)
	}
}
