package main

/**
 *	勉強用にクロールするもの作ってみた
 * 	今後はこれをDBに入れるところまで作ってみようかな
 */

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

//構造体(ハッシュみたいなもん)
type Result struct {
	Url  string
	star string
}

//定数
const (
	PAGE_LIMIT = 1000
)

var results = []Result{}

// チャネル
var finished = make(chan bool)

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

	finished <- true

}

func main() {
	url1 := "http://jobtalk.jp/company/index_1.html"
	url2 := "http://jobtalk.jp/company/index_2.html"

	go GetPage(url1)
	go GetPage(url2)

	// 終わるまで待つ
	<-finished
	<-finished

	for _, result := range results {
		fmt.Println(result.Url)
		fmt.Println(result.star)
	}
}
