package main

import (
     "fmt"
     "github.com/PuerkitoBio/goquery"
)

type Result struct {
     Url string
     star string
}


func GetPage(url string) []Result {

	 results := []Result{}
	 var result Result
     doc, _ := goquery.NewDocument(url)
     doc.Find("div.listBox").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.ChildrenFiltered("div.ttl").Find("a").Attr("href")
          star := s.ChildrenFiltered("p.point").Find("em").Text()
          //star := 4.0
          

          result.Url = url
          result.star = star
          results = append(results, result)
     })
     return results
}

func main() {
     url := "http://jobtalk.jp/company/index_1.html"
     results := GetPage(url)
     for _, result := range results {
          fmt.Println(result.Url)
          fmt.Println(result.star)
     }
}